package template

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
	"github.com/stankomichal/templie/internal/contextKey"
	"github.com/stankomichal/templie/internal/helpers"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"sort"
)

type TemplateHandler struct {
	templates map[string]Template `yaml:"templates"`
	config    *config.Config
}

func DefaultTemplateHandler() *TemplateHandler {
	return &TemplateHandler{
		templates: make(map[string]Template),
	}
}

func (th *TemplateHandler) SetConfig(config *config.Config) {
	th.config = config
}

func (th *TemplateHandler) InitializeTemplate(cmd *cobra.Command, templateName string, categories *[]string, copyContent bool) (string, error) {
	if _, exists := th.templates[templateName]; exists {
		return "", fmt.Errorf("template \"%s\" already exists", templateName)
	}

	// Default template
	template := Template{
		Name:       templateName,
		Categories: []string{th.config.DefaultCategory},
	}

	if categories != nil && len(*categories) != 0 {
		template.Categories = *categories
	}

	th.templates[templateName] = template

	if err := writeTemplateFile(th); err != nil {
		return "", fmt.Errorf("could not write template file: %w", err)
	}

	templateDir := filepath.Join(th.config.TemplateFolder, templateName)

	if err := os.MkdirAll(templateDir, 0755); err != nil {
		return "", fmt.Errorf("could not create template directory: %w", err)
	}

	if copyContent {
		currentDir, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("could not get current working directory: %w", err)
		}
		if err = copyDir(cmd, currentDir, templateDir); err != nil {
			return "", fmt.Errorf("could not copy working directory: %w", err)
		}
	}

	return templateDir, nil
}

func (th *TemplateHandler) CreateTemplate(cmd *cobra.Command, templateName string, path string, force bool) (*Template, error) {
	template, exists := th.templates[templateName]
	if !exists {
		return nil, fmt.Errorf("template \"%s\" does not exist", templateName)
	}

	templatePath := filepath.Join(th.config.TemplateFolder, templateName)

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("template \"%s\" does not exist", templateName)
	}

	if force {
		ctx := context.WithValue(cmd.Context(), contextKey.OverwriteKey, true)
		cmd.SetContext(ctx)
	}

	if err := copyDir(cmd, templatePath, path); err != nil {
		return nil, fmt.Errorf("could not copy directory: %w", err)
	}

	return &template, nil
}

func (th *TemplateHandler) GetTemplates() []Template {
	templates := make([]Template, 0, len(th.templates))
	for _, template := range th.templates {
		templates = append(templates, template)
	}
	sort.Slice(templates, func(i, j int) bool {
		return templates[i].Name < templates[j].Name
	})
	return templates
}

func (th *TemplateHandler) GetTemplate(templateName string) (*Template, error) {
	template, exists := th.templates[templateName]
	if !exists {
		return nil, fmt.Errorf("template \"%s\" does not exist", templateName)
	}
	return &template, nil
}

func (th *TemplateHandler) RenameTemplate(oldTemplateName string, newTemplateName string) (*Template, error) {
	template, exists := th.templates[oldTemplateName]
	if !exists {
		return nil, fmt.Errorf("template \"%s\" does not exist", oldTemplateName)
	}
	if _, exists = th.templates[newTemplateName]; exists {
		return nil, fmt.Errorf("template \"%s\" already exists", newTemplateName)
	}
	template.Name = newTemplateName
	th.templates[newTemplateName] = template
	delete(th.templates, oldTemplateName)
	if err := writeTemplateFile(th); err != nil {
		return nil, fmt.Errorf("could not write template file: %w", err)
	}
	templateDir := filepath.Join(th.config.TemplateFolder, oldTemplateName)
	newTemplateDir := filepath.Join(th.config.TemplateFolder, newTemplateName)
	if err := os.Rename(templateDir, newTemplateDir); err != nil {
		return nil, fmt.Errorf("could not rename template directory: %w", err)
	}
	return &template, nil
}

func (th *TemplateHandler) RemoveTemplate(templateName string) error {
	if _, exists := th.templates[templateName]; !exists {
		return fmt.Errorf("template \"%s\" does not exist", templateName)
	}
	delete(th.templates, templateName)
	if err := writeTemplateFile(th); err != nil {
		return fmt.Errorf("could not write template file: %w", err)
	}
	templateDir := filepath.Join(th.config.TemplateFolder, templateName)
	if err := os.RemoveAll(templateDir); err != nil {
		return fmt.Errorf("could not remove template directory: %w", err)
	}
	return nil
}

func (th *TemplateHandler) GetTemplatePath(templateName string) (string, error) {
	if _, exists := th.templates[templateName]; !exists {
		return "", fmt.Errorf("template \"%s\" does not exist", templateName)
	}

	templateDir := filepath.Join(th.config.TemplateFolder, templateName)
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		return "", fmt.Errorf("template \"%s\" does not exist", templateName)
	}
	return templateDir, nil
}

func (th *TemplateHandler) AddCategoryToTemplate(templateName string, category string) (*Template, error) {
	template, exists := th.templates[templateName]
	if !exists {
		return nil, fmt.Errorf("template \"%s\" does not exist", templateName)
	}
	for _, cat := range template.Categories {
		if cat == category {
			return nil, fmt.Errorf("category \"%s\" already exists in template \"%s\"", category, templateName)
		}
	}
	template.Categories = append(template.Categories, category)
	th.templates[templateName] = template
	if err := writeTemplateFile(th); err != nil {
		return nil, fmt.Errorf("could not write template file: %w", err)
	}
	return &template, nil
}

func (th *TemplateHandler) RemoveCategoryFromTemplate(templateName string, category string) ([]string, error) {
	template, exists := th.templates[templateName]
	if !exists {
		return nil, fmt.Errorf("template \"%s\" does not exist", templateName)
	}
	for i, cat := range template.Categories {
		if cat == category {
			template.Categories = append(template.Categories[:i], template.Categories[i+1:]...)
			th.templates[templateName] = template
			if err := writeTemplateFile(th); err != nil {
				return nil, fmt.Errorf("could not write template file: %w", err)
			}
			return template.Categories, nil
		}
	}
	return nil, fmt.Errorf("category \"%s\" does not exist in template \"%s\"", category, templateName)
}

func (th *TemplateHandler) SyncTemplates(syncStrategy string) error {
	switch syncStrategy {
	case "create":
		return th.syncByCreate()
	case "clean":
		return th.syncByClean()
	default:
		return fmt.Errorf("invalid sync strategy: %s", syncStrategy)
	}
}

func (th *TemplateHandler) syncByCreate() error {
	templateFolder := th.config.TemplateFolder
	templates := th.GetTemplates()

	dirs, err := os.ReadDir(templateFolder)
	if err != nil {
		return fmt.Errorf("could not read template folder: %w", err)
	}
	existingDirs := make(map[string]bool)
	for _, dir := range dirs {
		if dir.IsDir() {
			existingDirs[dir.Name()] = true
		}
	}

	existingMeta := make(map[string]bool)
	for _, template := range templates {
		existingMeta[template.Name] = true
		if !existingDirs[template.Name] {
			templatePath := filepath.Join(templateFolder, template.Name)
			if err = os.MkdirAll(templatePath, 0755); err != nil {
				return fmt.Errorf("could not create template directory: %w", err)
			}
		}
	}

	for dir := range existingDirs {
		if !existingMeta[dir] {
			templates = append(templates, Template{
				Name:       dir,
				Categories: []string{th.config.DefaultCategory},
			})
		}
	}

	th.templates = make(map[string]Template)
	for _, template := range templates {
		th.templates[template.Name] = template
	}

	if err = writeTemplateFile(th); err != nil {
		return err
	}

	return nil
}

func (th *TemplateHandler) syncByClean() error {
	templateFolder := th.config.TemplateFolder
	templates := th.GetTemplates()

	dirs, err := os.ReadDir(templateFolder)
	if err != nil {
		return fmt.Errorf("could not read template folder: %w", err)
	}

	existingDirs := make(map[string]bool)
	for _, dir := range dirs {
		if dir.IsDir() {
			existingDirs[dir.Name()] = true
		}
	}

	newTemplates := make([]Template, 0)
	for _, template := range templates {
		if existingDirs[template.Name] {
			newTemplates = append(newTemplates, template)
		}
	}

	existingMeta := make(map[string]bool)
	for _, template := range templates {
		existingMeta[template.Name] = true
	}

	for dir := range existingDirs {
		if !existingMeta[dir] {
			templatePath := filepath.Join(templateFolder, dir)
			if err = os.RemoveAll(templatePath); err != nil {
				return fmt.Errorf("could not remove template directory: %w", err)
			}
		}
	}
	th.templates = make(map[string]Template)
	for _, template := range newTemplates {
		th.templates[template.Name] = template
	}

	if err = writeTemplateFile(th); err != nil {
		return err
	}
	return nil
}

func Load(cmd *cobra.Command) (*TemplateHandler, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get home dir: %w", err)
	}
	templateFile := filepath.Join(homeDir, ".config", "templie", "templates.yaml")

	// Check if the template file exists
	if _, err = os.Stat(templateFile); os.IsNotExist(err) {
		templateHandler := DefaultTemplateHandler()

		if err = writeTemplateFile(templateHandler); err != nil {
			return nil, fmt.Errorf("could not write template file: %w", err)
		}

		cmd.Printf("Template file created at: %s\n", templateFile)
	}

	data, err := os.ReadFile(templateFile)
	if err != nil {
		return nil, fmt.Errorf("could not read template file: %w", err)
	}

	var templates map[string]Template
	if err = yaml.Unmarshal(data, &templates); err != nil {
		return nil, fmt.Errorf("could not parse template file: %w", err)
	}
	templateHandler := TemplateHandler{
		templates: templates,
	}

	// Templates are empty - default to an empty map
	if templateHandler.templates == nil {
		templateHandler.templates = make(map[string]Template)
	}

	return &templateHandler, nil
}

func writeTemplateFile(handler *TemplateHandler) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get home dir: %w", err)
	}

	templateFile := filepath.Join(homeDir, ".config", "templie", "templates.yaml")

	out, err := yaml.Marshal(handler.templates)

	if err != nil {
		return fmt.Errorf("could not marshal template handler: %w", err)
	}

	if err = os.WriteFile(templateFile, out, 0644); err != nil {
		return fmt.Errorf("could not write template file: %w", err)
	}

	return nil
}

func copyDir(cmd *cobra.Command, src string, dest string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("could not stat source directory: %w", err)
	}

	if !srcInfo.IsDir() {
		return fmt.Errorf("source is not a directory")
	}
	if err = os.MkdirAll(dest, srcInfo.Mode()); err != nil {
		return fmt.Errorf("could not create destination directory: %w", err)
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("could not read source directory: %w", err)
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			if err := copyDir(cmd, srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(cmd, srcPath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(cmd *cobra.Command, src string, dest string) error {
	if _, err := os.Stat(src); err == nil {
		overwrite := cmd.Context().Value(contextKey.OverwriteKey)
		if overwrite != nil {
			if !(overwrite.(bool)) {
				return nil
			}
		} else {
			res, err := helpers.ConfirmOverwrite(cmd, src)
			if err != nil {
				return fmt.Errorf("could not confirm overwrite: %w", err)
			}

			switch res {
			case helpers.ResponseNo:
				return nil
			case helpers.ResponseNoToAll:
				ctx := context.WithValue(cmd.Context(), contextKey.OverwriteKey, false)
				cmd.SetContext(ctx)
				return nil
			case helpers.ResponseYesToAll:
				ctx := context.WithValue(cmd.Context(), contextKey.OverwriteKey, true)
				cmd.SetContext(ctx)
			case helpers.ResponseYes:
				// Do nothing, just continue
			default:
				return fmt.Errorf("invalid response: %s", res)
			}
		}
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open source file: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("could not create destination file: %w", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}
