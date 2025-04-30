/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/stankomichal/templie/internal/helpers"
	"github.com/stankomichal/templie/internal/template"

	"github.com/spf13/cobra"
)

var categories []string
var copyContent bool

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <template-name>",
	Short: "Initialize a new template with optional category and file copying",
	Long: `
Initializes a new template by creating its metadata and associated directory.

You must provide a name for the template. The name will be sanitized, allowing only alphanumeric characters, underscores (_), dashes (-), and dots (.).

Optionally, you can use flags to assign categories and copy the contents of the current working directory 
into the newly created template folder.

Examples:
  templie template init my-template
  templie template init my-template --categories dev backend
  templie template init my-template --copy
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateHandler := cmd.Context().Value("templateHandler").(*template.TemplateHandler)

		templateName := helpers.SanitizeName(args[0])
		if templateName == "" {
			cmd.Println("Error: Template name after sanitization is empty. Valid characters are a-z, A-Z, 0-9, _, . and -")
			return
		}

		path, err := templateHandler.InitializeTemplate(templateName, &categories, copyContent)
		if err != nil {
			cmd.Printf("Error initializing template: %v\n", err)
			return
		}

		cmd.Printf("Template initialized at: %s\n", path)
	},
}

func init() {
	TemplateCmd.AddCommand(initCmd)

	initCmd.Flags().StringSliceVarP(&categories, "categories", "c", []string{}, "Categories for the template")
	initCmd.Flags().BoolVar(&copyContent, "copy-content", false, "Initialize template from current working directory")
}
