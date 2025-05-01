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
		templateHandler := cmd.Context().Value(contextKey.TemplateHandlerKey).(*template.TemplateHandler)
		oldName := helpers.SanitizeName(args[0])
		newName := helpers.SanitizeName(args[1])

		if oldName == "" {
			cmd.Println("Error: Old template name after sanitization is empty. Valid characters are a-z, A-Z, 0-9, _, . and -")
			return
		}
		if newName == "" {
			cmd.Println("Error: New template name after sanitization is empty. Valid characters are a-z, A-Z, 0-9, _, . and -")
			return
		}

		if _, err := templateHandler.RenameTemplate(oldName, newName); err != nil {
			cmd.Println("Error renaming template:", err)
			return
		}
		cmd.Printf("Template %s renamed to %s\n", oldName, newName)
	},
}

func init() {
	TemplateCmd.AddCommand(renameCmd)
}
