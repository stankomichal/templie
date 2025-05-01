/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package config

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
	"github.com/stankomichal/templie/internal/contextKey"
	"github.com/stankomichal/templie/internal/helpers"
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
		ctx := cmd.Context()
		cfg := ctx.Value(contextKey.ConfigKey).(*config.Config)

		helpers.VerbosePrintln(ctx, "Starting display config process")

		helpers.VerbosePrintln(ctx, "Marshaling config to YAML")
		out, err := yaml.Marshal(cfg)
		if err != nil {
			cmd.PrintErrf("Error marshaling config: %v\n", err)
			return
		}
		helpers.VerbosePrintln(ctx, "Config marshaled successfully")
		cmd.Print(string(out))

		helpers.VerbosePrintln(ctx, "Display config process completed")
	},
}

func init() {
	ConfigCmd.AddCommand(displayCmd)
}
