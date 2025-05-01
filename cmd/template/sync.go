/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package template

import (
	"github.com/stankomichal/templie/internal/contextKey"
	"github.com/stankomichal/templie/internal/helpers"
	"github.com/stankomichal/templie/internal/template"

	"github.com/spf13/cobra"
)

var syncStrategy string

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize templates based on the given strategy",
	Long: `
Synchronizes templates according to the specified strategy.
You can choose between two strategies: 
  - "create": Create any missing templates. (Default)
  - "clean": Clean up any templates that are no longer needed.

Examples:
  templie sync --strategy       // Synchronizes templates by creating missing ones
  templie sync --strategy clean // Cleans up unnecessary templates
`,

	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"create", "clean"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		helpers.VerbosePrintln(ctx, "Starting template sync process")
		helpers.VerbosePrintf(ctx, "Sync strategy: %s\n", syncStrategy)

		// Handle wrong strategy
		if syncStrategy != "create" && syncStrategy != "clean" {
			cmd.PrintErrln("Invalid strategy. Use 'create' or 'clean'")
			err := cmd.Help()
			if err != nil {
				helpers.VerbosePrintf(ctx, "Error displaying help: %v\n", err)
				return
			}

			cmd.SilenceErrors = true
			cmd.SilenceUsage = true
			return
		}

		templateHandler := ctx.Value(contextKey.TemplateHandlerKey).(*template.TemplateHandler)

		helpers.VerbosePrintf(ctx, "Syncing templates with strategy: %s\n", syncStrategy)
		err := templateHandler.SyncTemplates(syncStrategy)
		if err != nil {
			cmd.PrintErrf("Error syncing templates: %v\n", err)
			return
		}

		cmd.Println("Templates synchronized successfully")
		helpers.VerbosePrintln(ctx, "Template sync process completed successfully")
	},
}

func init() {
	TemplateCmd.AddCommand(syncCmd)
	syncCmd.Flags().StringVarP(&syncStrategy, "strategy", "s", "create", "Sync strategy: create or clean")
}
