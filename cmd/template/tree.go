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

var useIcons bool
var useColor bool

// treeCmd represents the tree command
var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Get the file tree of a template",
	Long: `
Displays the file and folder structure of a specific template.

If no template name is provided, an interactive menu will let you select one.

You can customize the tree with flags such as --icons or --color.

Examples:
  templie template tree
  templie template tree my-template
  templie template tree --icons --color
`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		templateHandler := ctx.Value(contextKey.TemplateHandlerKey).(*template.TemplateHandler)

		helpers.VerbosePrintln(cmd, "Starting template tree process")
		helpers.VerbosePrintf(cmd, "Using icons: %v, Using color: %v\n", useIcons, useColor)

		var templateName string
		if len(args) == 0 {
			helpers.VerbosePrintln(cmd, "No template name provided, prompting for selection")
			selected, err := template.SelectTemplateWithCategories(templateHandler.GetTemplates())
			if err != nil {
				cmd.PrintErrf("Error selecting template: %v\n", err)
				return
			}
			templateName = selected
		} else {
			templateName = args[0]
			helpers.VerbosePrintf(cmd, "Template name provided: %s\n", templateName)
		}

		helpers.VerbosePrintf(cmd, "Retrieving path for template: %s\n", templateName)
		templatePath, err := templateHandler.GetTemplatePath(templateName)
		if err != nil {
			cmd.PrintErrf("Error getting template: %v\n", err)
			return
		}

		helpers.VerbosePrintf(cmd, "Template path: %s\n", templatePath)
		cmd.Printf("Tree of template: %s\n", templateName)

		helpers.VerbosePrintln(cmd, "Generating tree preview")
		if err := helpers.PrintTree(cmd, templatePath, "", useIcons, useColor); err != nil {
			cmd.PrintErrf("Error printing template tree: %v\n", err)
			return
		}

		helpers.VerbosePrintln(cmd, "Template tree process completed")
	},
}

func init() {
	TemplateCmd.AddCommand(treeCmd)
	treeCmd.Flags().BoolVarP(&useIcons, "icons", "i", false, "Use icons in the tree")
	treeCmd.Flags().BoolVarP(&useColor, "color", "c", false, "Use color in the tree")
}
