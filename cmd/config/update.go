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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <variable-name> <new-value>",
	Short: "Update the value of a configuration variable",
	Long: `
The 'update' command allows you to update the value of a specific configuration variable.

This command requires the name of the configuration variable and the new value you want to set for it.

Examples:
  templie config update <variable-name> <new-value>   // Updates the specified configuration variable with the new value
`,

	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		cfg := ctx.Value(contextKey.ConfigKey).(*config.Config)

		helpers.VerbosePrintln(cmd, "Starting config update process")
		helpers.VerbosePrintf(cmd, "Variable name: %s\n", args[0])
		helpers.VerbosePrintf(cmd, "New value: %s\n", args[1])

		helpers.VerbosePrintf(cmd, "Updating variable %s to value %s\n", args[0], args[1])
		varValue, err := cfg.Update(args[0], args[1])
		if err != nil {
			cmd.PrintErrf("Error updating variable: %v\n", err)
			return
		}
		helpers.VerbosePrintf(cmd, "Variable updated to %s\n", varValue)
		cmd.Println("Variable updated successfully")

		helpers.VerbosePrintln(cmd, "Config update process completed")
	},
}

func init() {
	ConfigCmd.AddCommand(updateCmd)
}
