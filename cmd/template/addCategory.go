/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/helpers"
	"github.com/stankomichal/templie/internal/template"
)

// addCategoryCmd represents the addCategory command
var addCategoryCmd = &cobra.Command{
	Use:   "add-category <template-name> [categories...]",
	Short: "A brief description of your command",

	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		templateHandler := cmd.Context().Value("templateHandler").(*template.TemplateHandler)

		templateName := args[0]
		categories := args[1:]
		for _, category := range categories {

			sanitizedCategory := helpers.SanitizeName(category)
			if sanitizedCategory == "" {
				cmd.Println("Error: Category name after sanitization is empty. Valid characters are a-z, A-Z, 0-9, _, . and -")
				return
			}

			if _, err := templateHandler.AddCategoryToTemplate(templateName, sanitizedCategory); err != nil {
				cmd.Println("Error adding category:", err)
				return
			}
		}
		cmd.Printf("Categories %v added to template %s\n", categories, templateName)
	},
}

func init() {
	TemplateCmd.AddCommand(addCategoryCmd)
}
