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

var raw bool

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show <variable-name>",
	Short: "Show the current value of a configuration variable",
	Long: `
The 'show' command allows you to view the current value of a specific configuration variable.

This command takes the name of the configuration variable as an argument and outputs its current value.

Examples:
  templie config show <variable-name>   // Displays the current value of the specified configuration variable
`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		cfg := ctx.Value(contextKey.ConfigKey).(*config.Config)

		helpers.VerbosePrintln(cmd, ctx, "Starting show config variable process")
		helpers.VerbosePrintf(cmd, ctx, "Variable name: %s\n", args[0])

		varValue, err := cfg.Show(args[0])
		if err != nil {
			cmd.PrintErrf("Error showing variable: %v\n", err)
			return
		}
		helpers.VerbosePrintf(cmd, ctx, "Variable value retrieved successfully\n")
		if raw {
			cmd.Println(varValue)
			return
		} else {
			cmd.Printf("%s=%s\n", args[0], varValue)
		}

		helpers.VerbosePrintln(cmd, ctx, "Show config variable process completed")
	},
}

func init() {
	ConfigCmd.AddCommand(showCmd)

	showCmd.Flags().BoolVarP(&raw, "raw", "r", false, "Output only the raw value, useful for scripts")
}
