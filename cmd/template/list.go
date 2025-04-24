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
	Short: "A brief description of your command",

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
