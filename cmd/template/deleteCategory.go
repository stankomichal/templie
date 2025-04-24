/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/stankomichal/templie/internal/template"

	"github.com/spf13/cobra"
)

// deleteCategoryCmd represents the deleteCategory command
var deleteCategoryCmd = &cobra.Command{
	Use:   "delete-category <template-name> [categories...]>",
	Short: "A brief description of your command",

	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		templateHandler := cmd.Context().Value("templateHandler").(*template.TemplateHandler)

		templateName := args[0]
		categories := args[1:]
		for _, category := range categories {
			if _, err := templateHandler.DeleteCategoryFromTemplate(templateName, category); err != nil {
				cmd.Println("Error deleting category:", err)
				return
			}
		}
		cmd.Printf("Categories %v deleted from template %s\n", categories, templateName)
	},
}

func init() {
	TemplateCmd.AddCommand(deleteCategoryCmd)
}
