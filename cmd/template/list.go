/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/stankomichal/templie/internal/template"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved templates and their categories",
	Long: `
Displays a list of all available templates along with their associated categories.

Each template entry will show its name and a nested list of any categories it belongs to.

Examples:
  templie template list
`,

	Run: func(cmd *cobra.Command, args []string) {
		templateHandler := cmd.Context().Value("templateHandler").(*template.TemplateHandler)

		templates := templateHandler.GetTemplates()

		if len(templates) == 0 {
			cmd.Println("No templates found.")
			return
		}
		cmd.Println("Templates:")
		for _, template := range templates {
			cmd.Printf("- %s\n", template.Name)
			for _, category := range template.Categories {
				cmd.Printf("\t- %s\n", category)
			}
		}
	},
}

func init() {
	TemplateCmd.AddCommand(listCmd)
}
