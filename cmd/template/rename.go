/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/helpers"
	"github.com/stankomichal/templie/internal/template"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "A brief description of your command",

	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		templateHandler := cmd.Context().Value("templateHandler").(*template.TemplateHandler)
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
