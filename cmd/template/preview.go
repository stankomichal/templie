/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
	"github.com/stankomichal/templie/internal/helpers"
	"github.com/stankomichal/templie/internal/template"
)

var useIcons bool
var useColor bool

// previewCmd represents the preview command
var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "A brief description of your command",

	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := cmd.Context().Value("config").(*config.Config)
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

		templatePath, err := templateHandler.GetTemplatePath(templateName)
		if err != nil {
			cmd.Println("Error getting template:", err)
			return
		}
		cmd.Println("Preview of template:", templateName)
		if err := helpers.PrintTree(cfg.FileDecorators, templatePath, "", useIcons, useColor); err != nil {
			cmd.Println("Error printing template tree:", err)
			return
		}
	},
}

func init() {
	TemplateCmd.AddCommand(previewCmd)
	previewCmd.Flags().BoolVarP(&useIcons, "icons", "i", false, "Use icons in the preview")
	previewCmd.Flags().BoolVarP(&useColor, "color", "c", false, "Use color in the preview")
}
