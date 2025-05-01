/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package config

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
	"github.com/stankomichal/templie/internal/contextKey"
	"github.com/stankomichal/templie/internal/helpers"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset <variable-name>",
	Short: "Reset a specific configuration variable to its default value",
	Long: `
The 'reset' command allows you to reset a specific configuration variable to its default value.

This command accepts the name of the configuration variable you want to reset and reverts it to its predefined default value.

Examples:
  templie config reset <variable-name>   // Resets the specified configuration variable to its default
`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		cfg := ctx.Value(contextKey.ConfigKey).(*config.Config)

		helpers.VerbosePrintln(ctx, "Starting config reset process")
		helpers.VerbosePrintf(ctx, "Variable name: %s\n", args[0])

		helpers.VerbosePrintf(ctx, "Resetting variable %s to default value\n", args[0])
		varValue, err := cfg.Reset(args[0])
		if err != nil {
			cmd.PrintErrf("Error resetting variable: %v\n", err)
			return
		}
		cmd.Printf("Variable reset successfully to: %s\n", varValue)

		helpers.VerbosePrintln(ctx, "Config reset process completed")
	},
}

func init() {
	ConfigCmd.AddCommand(resetCmd)
}
