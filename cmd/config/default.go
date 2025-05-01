/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package config

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
	"github.com/stankomichal/templie/internal/helpers"
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
		ctx := cmd.Context()

		helpers.VerbosePrintln(ctx, "Starting display default config process")

		helpers.VerbosePrintln(ctx, "Getting default config")
		defaultConfig := config.DefaultConfig()

		helpers.VerbosePrintln(ctx, "Marshaling default config to YAML")
		out, err := yaml.Marshal(defaultConfig)

		if err != nil {
			cmd.PrintErrf("Error marshaling config: %v\n", err)
			return
		}
		helpers.VerbosePrintln(ctx, "Default config marshaled successfully")
		cmd.Print(string(out))

		helpers.VerbosePrintln(ctx, "Display default config process completed")
	},
}

func init() {
	ConfigCmd.AddCommand(defaultCmd)
}
