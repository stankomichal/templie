package helpers

import (
	"github.com/fatih/color"
	"github.com/stankomichal/templie/internal/config"
	"path/filepath"
)

func GetFileIconAndColor(fileDecorators map[string]config.FileDecorator, fileName string, useIcons bool, useColor bool) (string, *color.Color) {
	fileIcon := ""
	fileColor := color.New(color.Reset)
	color.New()
	if useIcons {
		if foundIcon, exists := fileDecorators[filepath.Ext(fileName)]; exists {
			fileIcon = foundIcon.Icon
		} else {
			fileIcon = "📄"
		}
	}
	if useColor {
		if foundColor, exists := fileDecorators[filepath.Ext(fileName)]; exists {
			fileColor = foundColor.Color
		} else {
			fileColor = color.New(color.FgGreen)
		}
	}

	return fileIcon + " ", fileColor
}
