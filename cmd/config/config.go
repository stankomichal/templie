/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package config

import (
	"github.com/spf13/cobra"
)

// ConfigCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings for your templates",
	Long: `
The 'config' command allows you to manage configuration settings for your templates, 
including viewing and modifying settings related to template categories, file paths, and other preferences.

Use this command to:
  - View current configuration settings
  - Update configuration settings for templates
  - Reset configuration to default settings

Examples:
  templie config display                                // Displays the current configuration
  templie config show <variable-name>                   // Shows the value of a specific configuration key
  templie config update <variable-name> <new-value>  	// Updates the configuration setting for a specific key
  templie config reset <variable-name>                  // Resets all configuration settings to their defaults
`,
}

func init() {

}
