/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
	"github.com/stankomichal/templie/internal/contextKey"
	"github.com/stankomichal/templie/internal/helpers"
	"github.com/stankomichal/templie/internal/template"
)

var useIcons bool
var useColor bool

// previewCmd represents the preview command
var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Show the file tree of a template",
	Long: `
Displays a preview of the file and folder structure of a specific template.

If no template name is provided, an interactive menu will let you select one.

You can customize the preview with flags such as --icons or --color.

Examples:
  templie template preview my-template
  templie template preview --icons --color
  templie template preview
`,

	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		cfg := ctx.Value(contextKey.ConfigKey).(*config.Config)
		templateHandler := ctx.Value(contextKey.TemplateHandlerKey).(*template.TemplateHandler)

		helpers.VerbosePrintln(ctx, "Starting template preview process")
		helpers.VerbosePrintf(ctx, "Using icons: %v, Using color: %v\n", useIcons, useColor)

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

		helpers.VerbosePrintf(ctx, "Retrieving path for template: %s\n", templateName)
		templatePath, err := templateHandler.GetTemplatePath(templateName)
		if err != nil {
			cmd.PrintErrf("Error getting template: %v\n", err)
			return
		}

		helpers.VerbosePrintf(ctx, "Template path: %s\n", templatePath)
		cmd.Printf("Preview of template: %s\n", templateName)

		helpers.VerbosePrintln(ctx, "Generating tree preview")
		if err := helpers.PrintTree(cfg.FileDecorators, templatePath, "", useIcons, useColor); err != nil {
			cmd.PrintErrf("Error printing template tree: %v\n", err)
			return
		}

		helpers.VerbosePrintln(ctx, "Template preview process completed")
	},
}

func init() {
	TemplateCmd.AddCommand(previewCmd)
	previewCmd.Flags().BoolVarP(&useIcons, "icons", "i", false, "Use icons in the preview")
	previewCmd.Flags().BoolVarP(&useColor, "color", "c", false, "Use color in the preview")
}
