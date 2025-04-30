/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package config

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
	"gopkg.in/yaml.v3"
)

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "Display the current configuration",
	Long: `
The 'display' command outputs the current configuration settings.

This command retrieves and displays the active configuration settings, formatted as a YAML file.

Examples:
  templie config display   // Displays the current configuration in YAML format
`,

	Run: func(cmd *cobra.Command, args []string) {
		config := cmd.Context().Value("config").(*config.Config)

		out, err := yaml.Marshal(config)
		if err != nil {
			cmd.Printf("could not marshal config: %v\n", err)
			return
		}

		cmd.Print(string(out))
	},
}

func init() {
	ConfigCmd.AddCommand(displayCmd)
}
