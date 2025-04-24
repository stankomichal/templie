/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/template"
)

// pathCmd represents the path command
var pathCmd = &cobra.Command{
	Use:   "path",
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
		path, err := templateHandler.GetTemplatePath(templateName)
		if err != nil {
			cmd.Println("Error getting template path:", err)
			return
		}
		cmd.Printf("Path for template %s: %s\n", templateName, path)
	},
}

func init() {
	TemplateCmd.AddCommand(pathCmd)
}
