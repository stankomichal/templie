/*
Copyright © 2025 Michal Stanko michal.stankoml@gmail.com
*/
package cmd

import (
	"context"
	configCmd "github.com/stankomichal/templie/cmd/config"
	templateCmd "github.com/stankomichal/templie/cmd/template"
	configClass "github.com/stankomichal/templie/internal/config"
	"github.com/stankomichal/templie/internal/template"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "templie",
	Short: "A simple and flexible CLI tool for managing project templates",
	Long: `
Templie is a cross-platform CLI tool designed to streamline the management, creation,
and initialization of reusable project templates.

With templie, you can:
  - Create and organize templates in categories
  - Initialize new projects from templates
  - Sync metadata with template folders
  - Add icons and colorized tree views
  - Manage templates through a clean YAML-based config

Whether you're a solo developer or managing shared boilerplate across teams,
templie helps you keep your templates consistent, structured, and accessible.


Examples:
  templie template init my-template
  templie template create my-template -o output-folder
  templie template preview my-template -ci
  templie template sync --strategy=create
  templie config display
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Load the config file
		cfg, err := configClass.Load()
		if err != nil {
			cmd.Println("Failed to load config:", err)
		}
		ctx := context.WithValue(cmd.Context(), "config", cfg)

		// Load the template file
		templateHandler, err := template.Load()
		if err != nil {
			cmd.Println("Failed to load templates:", err)
		}
		templateHandler.SetConfig(cfg)
		ctx = context.WithValue(ctx, "templateHandler", templateHandler)
		cmd.SetContext(ctx)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(configCmd.ConfigCmd)
	rootCmd.AddCommand(templateCmd.TemplateCmd)
}

func init() {
	rootCmd.Version = "0.1.0"
	rootCmd.SetVersionTemplate("{{printf \"templie version %s\" .Version}}\n")

	addSubcommandPalettes()
}
