/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/contextKey"
	"github.com/stankomichal/templie/internal/helpers"
	"github.com/stankomichal/templie/internal/template"
	"strings"
)

// addCategoryCmd represents the addCategory command
var addCategoryCmd = &cobra.Command{
	Use:   "add-category <template-name> <categories>",
	Short: "Add one or more categories to a template",
	Long: `
Adds one or more categories to the specified template.

Each category name will be sanitized to ensure it's a valid folder-safe name.
Valid characters include: a-z, A-Z, 0-9, _, . and -

Examples:
  templie template add-category my-template backend,shared
  templie template add-category "cool-template" devops

You must provide at least one category.
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		templateHandler := ctx.Value(contextKey.TemplateHandlerKey).(*template.TemplateHandler)

		helpers.VerbosePrintln(cmd, "Starting add category process")

		templateName := args[0]
		helpers.VerbosePrintf(cmd, "Template name: %s\n", templateName)

		categoriesString := args[1]
		categories := strings.Split(categoriesString, ",")
		helpers.VerbosePrintf(cmd, "Categories to add: %v\n", categories)

		for _, category := range categories {
			helpers.VerbosePrintf(cmd, "Processing category: %s\n", category)

			sanitizedCategory := helpers.SanitizeName(category)
			helpers.VerbosePrintf(cmd, "Sanitized category name: %s\n", sanitizedCategory)

			if sanitizedCategory == "" {
				cmd.PrintErrln("Error: Category name after sanitization is empty. Valid characters are a-z, A-Z, 0-9, _, . and -")
				return
			}

			helpers.VerbosePrintf(cmd, "Adding category %s to template %s\n", sanitizedCategory, templateName)
			if _, err := templateHandler.AddCategoryToTemplate(templateName, sanitizedCategory); err != nil {
				cmd.PrintErrf("Error adding category: %v\n", err)
				return
			}
			helpers.VerbosePrintf(cmd, "Successfully added category %s to template %s\n", sanitizedCategory, templateName)
		}

		cmd.Printf("Categories %v added to template %s\n", categories, templateName)
		helpers.VerbosePrintln(cmd, "Add category process completed successfully")
	},
}

func init() {
	TemplateCmd.AddCommand(addCategoryCmd)
}
