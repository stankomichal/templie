/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/template"
	"os"
)

var outputPath string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <template-name>",
	Short: "Create a new template folder with optional output folder",
	Long: `
Creates a new template by name. If no name is provided, you can interactively select from existing templates.

You can also use the --output flag to specify where the template folder should be created; otherwise, the current working directory is used.

Examples:
  templie template create my-template
  templie template create
  templie t c --output ./generated

If no name is provided, you’ll be prompted to choose from existing templates.
`,
	Aliases: []string{"c"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateHandler := cmd.Context().Value("templateHandler").(*template.TemplateHandler)
		if outputPath == "" {
			dir, err := os.Getwd()
			if err != nil {
				cmd.Printf("Error getting current directory: %v\n", err)
				return
			}
			outputPath = dir
		}

		var templateName string
		if len(args) == 0 {
			selected, err := template.SelectTemplateWithCategories(templateHandler.GetTemplates())
			if err != nil {
				cmd.Println("Error selecting template:", err)
				return
			}
			templateName = selected
		} else {
			templateName = args[0]
		}

		if templateName == "" {
			cmd.Println("Error: Template name after sanitization is empty. Valid characters are a-z, A-Z, 0-9, _, . and -")
			return
		}

		_, err := templateHandler.CreateTemplate(templateName, outputPath)
		if err != nil {
			cmd.Printf("Error creating template: %v\n", err)
			return
		}
		cmd.Printf("Template successfully created.\n")
	},
}

func init() {
	TemplateCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output path for the template")
}
