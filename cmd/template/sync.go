/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/stankomichal/templie/internal/template"

	"github.com/spf13/cobra"
)

var syncStrategy string

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "A brief description of your command",

	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"create", "clean"},
	Run: func(cmd *cobra.Command, args []string) {
		// Handle wrong strategy
		if syncStrategy != "create" && syncStrategy != "clean" {
			cmd.PrintErrln("Invalid strategy. Use 'create' or 'clean'.")
			err := cmd.Help()
			if err != nil {
				return
			}
			cmd.SilenceErrors = true
			cmd.SilenceUsage = true
			return
		}

		templateHandler := cmd.Context().Value("templateHandler").(*template.TemplateHandler)
		err := templateHandler.SyncTemplates(syncStrategy)
		if err != nil {
			return
		}
	},
}

func init() {
	TemplateCmd.AddCommand(syncCmd)
	syncCmd.Flags().StringVarP(&syncStrategy, "strategy", "s", "create", "Sync strategy: create or clean")
}
