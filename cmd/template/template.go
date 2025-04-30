/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/spf13/cobra"
)

// TemplateCmd represents the template command
var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Manage and manipulate templates",
	Long: `
The 'template' command allows you to manage templates, including creating, 
modifying, deleting, and interacting with various template categories.

Use this command to:
  - Initialize templates or configure their categories
  - Create new templates
  - Add, remove or list categories for templates
  - Preview templates before use

Examples:
  templie template create <template-name>                          // Creates a new template with the given name
  templie template list                                            // Lists all templates with categories
  templie template remove <template-name>                          // Removes the specified template
  templie t add-category <template-name> [categories...]    // Adds categories to a template
  templie t remove-category <template-name> [categories...] // Removes categories from a template
`,
	Aliases: []string{"t"},
}

func init() {

}
