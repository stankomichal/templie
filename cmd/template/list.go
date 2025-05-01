/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/stankomichal/templie/internal/contextKey"
	"github.com/stankomichal/templie/internal/helpers"
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
  templie template ls
`,
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		templateHandler := ctx.Value(contextKey.TemplateHandlerKey).(*template.TemplateHandler)

		helpers.VerbosePrintln(cmd, "Starting template listing process")

		helpers.VerbosePrintln(cmd, "Retrieving templates from template handler")
		templates := templateHandler.GetTemplates()

		if len(templates) == 0 {
			cmd.Println("No templates found")
			return
		}

		helpers.VerbosePrintf(cmd, "Found %d templates\n", len(templates))
		cmd.Println("Templates:")
		for _, template := range templates {
			helpers.VerbosePrintf(cmd, "Processing template: %s\n", template.Name)
			cmd.Printf("- %s\n", template.Name)

			if len(template.Categories) > 0 {
				for _, category := range template.Categories {
					cmd.Printf("\t- %s\n", category)
				}
			} else {
				cmd.Printf("\t- No categories\n")
			}
		}

		helpers.VerbosePrintln(cmd, "Template listing process completed")
	},
}

func init() {
	TemplateCmd.AddCommand(listCmd)
}
