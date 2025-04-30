/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
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
		config := cmd.Context().Value("config").(*config.Config)
		varValue, err := config.Reset(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(varValue)
		}
	},
}

func init() {
	ConfigCmd.AddCommand(resetCmd)
}
