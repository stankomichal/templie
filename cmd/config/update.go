/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package config

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <variable-name> <new-value>",
	Short: "A brief description of your command",

	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("update called")
		config := cmd.Context().Value("config").(*config.Config)
		varValue, err := config.Update(args[0], args[1])
		if err != nil {
			cmd.Println(err)
		} else {
			cmd.Println(varValue)
		}
	},
}

func init() {
	ConfigCmd.AddCommand(updateCmd)
}
