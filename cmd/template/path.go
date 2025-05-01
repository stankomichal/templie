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

var raw bool

// pathCmd represents the path command
var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print the file system path of a specific template",
	Long: `Displays the full file system path to the directory of a given template.

If no template name is provided, an interactive selection menu will be shown.

Examples:
  templie template path my-template
  templie template path
`,

	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		templateHandler := ctx.Value(contextKey.TemplateHandlerKey).(*template.TemplateHandler)

		helpers.VerbosePrintln(cmd, "Starting template path retrieval process")

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
		path, err := templateHandler.GetTemplatePath(templateName)
		if err != nil {
			cmd.PrintErrf("Error getting template path: %v\n", err)
			return
		}
		if raw {
			cmd.Println(path)
		} else {
			cmd.Printf("Path for template %s: %s\n", templateName, path)
		}

		helpers.VerbosePrintln(cmd, "Template path retrieval process completed")
	},
}

func init() {
	TemplateCmd.AddCommand(pathCmd)

	pathCmd.Flags().BoolVarP(&raw, "raw", "r", false, "Output only the raw value, useful for scripts")
}
