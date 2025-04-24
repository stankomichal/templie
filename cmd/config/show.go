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
	Short: "A brief description of your command",

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
