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
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
