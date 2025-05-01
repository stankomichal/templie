/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/stankomichal/templie/internal/contextKey"
	"github.com/stankomichal/templie/internal/helpers"
	"github.com/stankomichal/templie/internal/template"
	"strings"

	"github.com/spf13/cobra"
)

// removeCategoryCmd represents the removeCategory command
var removeCategoryCmd = &cobra.Command{
	Use:   "remove-category <template-name> [categories...]>",
	Short: "Remove one or more categories from a template",
	Long: `
Removes one or more categories from the specified template.

You must provide the template name followed by at least one category to remove.

Examples:
  templie template remove-category my-template dev backend
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		templateHandler := ctx.Value(contextKey.TemplateHandlerKey).(*template.TemplateHandler)

		helpers.VerbosePrintln(ctx, "Starting category removal process")

		templateName := args[0]
		helpers.VerbosePrintf(ctx, "Template name: %s\n", templateName)

		categoriesString := args[1]
		categories := strings.Split(categoriesString, ",")
		helpers.VerbosePrintf(ctx, "Categories to remove: %v\n", categories)

		for _, category := range categories {
			helpers.VerbosePrintf(ctx, "Removing category %s from template %s\n", category, templateName)
			if _, err := templateHandler.RemoveCategoryFromTemplate(templateName, category); err != nil {
				cmd.PrintErrf("Error removing category: %v\n", err)
				return
			}
			helpers.VerbosePrintf(ctx, "Successfully removed category %s from template %s\n", category, templateName)
		}

		cmd.Printf("Categories %v removed from template %s\n", categories, templateName)
		helpers.VerbosePrintln(ctx, "Category removal process completed")
	},
}

func init() {
	TemplateCmd.AddCommand(removeCategoryCmd)
}
