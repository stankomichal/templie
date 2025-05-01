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
  - Set configuration settings for templates
  - Reset configuration to default settings

Examples:
  templie config display                                // Displays the current configuration
  templie config get <variable-name>                    // Shows the value of a specific configuration key
  templie config set <variable-name> <new-value>  	    // Updates the configuration setting of a specific key
  templie config reset <variable-name>                  // Resets all configuration settings to their defaults
`,
}

func init() {

}
