package template

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
)

func formatTemplateOptions(templates []Template) []string {
	var options []string

	for _, template := range templates {
		cat := "uncatogorized"
		if len(template.Categories) > 0 {
			cat = strings.Join(template.Categories, ", ")
		}

		label := fmt.Sprintf("%s [%s]", template.Name, cat)

		options = append(options, label)
	}

	return options
}

func SelectTemplateWithCategories(templates []Template) (string, error) {
	options := formatTemplateOptions(templates)

	prompt := promptui.Select{
		Label: "Select a template",
		Items: options,
		Size:  10,
		Searcher: func(input string, index int) bool {
			item := options[index]
			return strings.Contains(strings.ToLower(item), strings.ToLower(input))
		},
	}

	i, _, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed %v", err)
	}
	return templates[i].Name, nil
}
