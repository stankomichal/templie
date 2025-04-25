/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
)

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
		config := cmd.Context().Value("config").(*config.Config)
		varValue, err := config.Show(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(varValue)
		}
	},
}

func init() {
	ConfigCmd.AddCommand(showCmd)
}
