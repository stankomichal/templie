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

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the current configuration",
	Long: `
The 'view' command outputs the current configuration settings.

This command retrieves and displays the active configuration settings, formatted as a YAML file.

Examples:
  templie config view   // Displays the current configuration in YAML format
`,

	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		cfg := ctx.Value(contextKey.ConfigKey).(*config.Config)

		helpers.VerbosePrintln(cmd, "Starting view config process")

		helpers.VerbosePrintln(cmd, "Marshaling config to YAML")
		out, err := yaml.Marshal(cfg)
		if err != nil {
			cmd.PrintErrf("Error marshaling config: %v\n", err)
			return
		}
		helpers.VerbosePrintln(cmd, "Config marshaled successfully")
		cmd.Print(string(out))

		helpers.VerbosePrintln(cmd, "View config process completed")
	},
}

func init() {
	ConfigCmd.AddCommand(viewCmd)
}
