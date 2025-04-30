/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package config

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
	"gopkg.in/yaml.v3"
)

// defaultCmd represents the default command
var defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Display the default configuration",
	Long: `
The 'default' command displays the default configuration settings for the templates.

This command outputs the default configuration as a YAML file, which can be used as a template for setting up new template configurations.

Examples:
  templie config default   // Displays the default configuration settings in YAML format
`,

	Run: func(cmd *cobra.Command, args []string) {
		out, err := yaml.Marshal(config.DefaultConfig())

		if err != nil {
			cmd.Printf("could not marshal config: %v\n", err)
			return
		}

		cmd.Print(string(out))
	},
}

func init() {
	ConfigCmd.AddCommand(defaultCmd)
}
