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

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename an existing template",
	Long: `
Renames a template folder to a new name.

Both old and new names will be sanitized, allowing only alphanumeric characters, underscores (_), dashes (-), and dots (.).

Examples:
  templie template rename old-template new-template
  templie template rename my.template_1 my-renamed-template
`,

	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		templateHandler := ctx.Value(contextKey.TemplateHandlerKey).(*template.TemplateHandler)

		helpers.VerbosePrintln(cmd, "Starting template rename process")

		helpers.VerbosePrintf(cmd, "Raw old name: %s\n", args[0])
		oldName := helpers.SanitizeName(args[0])
		helpers.VerbosePrintf(cmd, "Sanitized old name: %s\n", oldName)

		helpers.VerbosePrintf(cmd, "Raw new name: %s\n", args[1])
		newName := helpers.SanitizeName(args[1])
		helpers.VerbosePrintf(cmd, "Sanitized new name: %s\n", newName)

		if oldName == "" {
			cmd.PrintErrln("Error: Old template name after sanitization is empty. Valid characters are a-z, A-Z, 0-9, _, . and -")
			return
		}
		if newName == "" {
			cmd.PrintErrln("Error: New template name after sanitization is empty. Valid characters are a-z, A-Z, 0-9, _, . and -")
			return
		}

		helpers.VerbosePrintf(cmd, "Renaming template from %s to %s\n", oldName, newName)
		if _, err := templateHandler.RenameTemplate(oldName, newName); err != nil {
			cmd.PrintErrf("Error renaming template: %v\n", err)
			return
		}

		cmd.Printf("Template %s successfully renamed to %s\n", oldName, newName)

		helpers.VerbosePrintln(cmd, "Template rename process completed")
	},
}

func init() {
	TemplateCmd.AddCommand(renameCmd)
}
