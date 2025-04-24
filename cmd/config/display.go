/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package config

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
	"gopkg.in/yaml.v3"
)

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		config := cmd.Context().Value("config").(*config.Config)

		out, err := yaml.Marshal(config)
		if err != nil {
			cmd.Printf("could not marshal config: %v\n", err)
			return
		}

		cmd.Print(string(out))
	},
}

func init() {
	ConfigCmd.AddCommand(displayCmd)
}
