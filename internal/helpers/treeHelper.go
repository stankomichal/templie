package helpers

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
	"os"
	"path/filepath"
)

func PrintTree(cmd *cobra.Command, fileDecorators map[string]config.FileDecorator, path string, prefix string, useIcons bool, useColor bool) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for i, entry := range entries {
		connector := "├── "
		if i == len(entries)-1 {
			connector = "└── "
		}

		if entry.IsDir() {
			dirIcon := ""
			if useIcons {
				dirIcon = "📁 "
			}
			dirColor := color.New(color.Reset)
			if useColor {
				dirColor = color.New(color.FgCyan, color.Bold)
			}
			cmd.Print(prefix + connector)
			_, _ = dirColor.Fprintln(cmd.OutOrStdout(), dirIcon+entry.Name())
			subPrefix := "│   "
			if i == len(entries)-1 {
				subPrefix = "    "
			}
			err = PrintTree(cmd, fileDecorators, filepath.Join(path, entry.Name()), prefix+subPrefix, useIcons, useColor)
			if err != nil {
				return err
			}
		} else {
			icon, fileColor := GetFileIconAndColor(fileDecorators, filepath.Join(path, entry.Name()), useIcons, useColor)
			cmd.Print(prefix + connector)
			_, _ = fileColor.Fprintln(cmd.OutOrStdout(), icon+entry.Name())
		}
	}
	return nil
}
