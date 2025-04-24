package helpers

import (
	"regexp"
	"strings"
)

func SanitizeName(name string) string {
	// Remove any leading or trailing whitespace
	name = strings.TrimSpace(name)

	// Replace spaces with underscores
	name = strings.ReplaceAll(name, " ", "_")

	// Remove any special characters (except underscores, dashes, and periods)
	re := regexp.MustCompile("[^a-z0-9-_.]")
	name = re.ReplaceAllString(name, "")

	return name
}
