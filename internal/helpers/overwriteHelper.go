package helpers

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"strings"
)

const (
	ResponseYes      = "y"
	ResponseNo       = "n"
	ResponseYesToAll = "ya"
	ResponseNoToAll  = "na"
)

func ConfirmOverwrite(cmd *cobra.Command, file string) (string, error) {
	for {
		cmd.Printf("File %s already exists. Do you want to overwrite it? [y]es / [ya] yes to all / [n]o / [na] no to all: ", file)
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil && err != io.EOF {
			return "", fmt.Errorf("failed to read input: %w", err)
		}

		response = strings.TrimSpace(strings.ToLower(response))
		switch response {
		case "y", "n", "ya", "na":
			return response, nil
		default:
			cmd.Println("Please enter 'y', 'n', 'ya' or 'na'.")
		}
	}
}
