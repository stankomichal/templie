/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/template"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove <template-name>",
	Short: "Remove a template and its associated metadata",
	Long: `
Removes a specified template and its metadata entry.

If no template name is provided, you will be prompted to select from the list of available templates.

Examples:
  templie template remove
  templie t rm my-template
`,
	Aliases: []string{"rm"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateHandler := cmd.Context().Value("templateHandler").(*template.TemplateHandler)

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

		if err := templateHandler.RemoveTemplate(templateName); err != nil {
			cmd.Println("Error removing the template:", err)
			return
		}
		cmd.Printf("Template %s removed\n", templateName)
	},
}

func init() {
	TemplateCmd.AddCommand(removeCmd)
}
