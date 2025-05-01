package helpers

import (
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/config"
	"github.com/stankomichal/templie/internal/contextKey"
	"os"
	"path/filepath"
)

func PrintTree(cmd *cobra.Command, path string, prefix string, useIcons bool, useColor bool) error {
	cfg := cmd.Context().Value(contextKey.ConfigKey).(*config.Config)
	fileDecorators := cfg.FileDecorators
	folderDecorator := cfg.FolderDecorator

	dirIcon, dirColor := GetFolderIconAndColor(folderDecorator, useIcons, useColor)

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

			cmd.Print(prefix + connector)
			_, _ = dirColor.Fprintln(cmd.OutOrStdout(), dirIcon+entry.Name())
			subPrefix := "│   "
			if i == len(entries)-1 {
				subPrefix = "    "
			}
			err = PrintTree(cmd, filepath.Join(path, entry.Name()), prefix+subPrefix, useIcons, useColor)
			if err != nil {
				return err
			}
		} else {
			fileIcon, fileColor := GetFileIconAndColor(fileDecorators, filepath.Ext(entry.Name()), useIcons, useColor)
			cmd.Print(prefix + connector)
			_, _ = fileColor.Fprintln(cmd.OutOrStdout(), fileIcon+entry.Name())
		}
	}
	return nil
}
