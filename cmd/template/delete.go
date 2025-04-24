/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/template"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <template-name>",
	Short: "A brief description of your command",

	Args: cobra.MaximumNArgs(1),
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

		if err := templateHandler.DeleteTemplate(templateName); err != nil {
			cmd.Println("Error deleting template:", err)
			return
		}
		cmd.Printf("Template %s deleted\n", templateName)
	},
}

func init() {
	TemplateCmd.AddCommand(deleteCmd)
}
