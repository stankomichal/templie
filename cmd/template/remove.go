/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/contextKey"
	"github.com/stankomichal/templie/internal/helpers"
	"github.com/stankomichal/templie/internal/template"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove <template-name>",
	Short: "Remove a template and its associated metadata",
	Long: `
Removes a specified template and its metadata entry.

If no template name is provided, you will be prompted to select from the list of available templates.

Examples:
  templie template remove
  templie t rm my-template
`,
	Aliases: []string{"rm"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		templateHandler := ctx.Value(contextKey.TemplateHandlerKey).(*template.TemplateHandler)

		helpers.VerbosePrintln(ctx, "Starting template removal process")

		var templateName string
		if len(args) == 0 {
			helpers.VerbosePrintln(ctx, "No template name provided, prompting for selection")
			selected, err := template.SelectTemplateWithCategories(templateHandler.GetTemplates())
			if err != nil {
				cmd.PrintErrf("Error selecting template: %v\n", err)
				return
			}
			templateName = selected
		} else {
			templateName = args[0]
			helpers.VerbosePrintf(ctx, "Template name provided: %s\n", templateName)
		}

		helpers.VerbosePrintf(ctx, "Removing template: %s\n", templateName)
		if err := templateHandler.RemoveTemplate(templateName); err != nil {
			cmd.PrintErrf("Error removing the template: %v\n", err)
			return
		}

		cmd.Printf("Template %s successfully removed\n", templateName)

		helpers.VerbosePrintln(ctx, "Template removal process completed")
	},
}

func init() {
	TemplateCmd.AddCommand(removeCmd)
}
