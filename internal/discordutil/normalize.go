package discordutil

import (
	"regexp"
	"strings"
)

var whitespaceRegex = regexp.MustCompile(`\s+`)

func NormalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = whitespaceRegex.ReplaceAllString(name, " ")
	name = whitespaceRegex.ReplaceAllString(name, "-")
	name = strings.ToLower(name)
	return name
}

func NormalizeNameConfig(name string) string {
	name = strings.TrimSpace(name)
	name = whitespaceRegex.ReplaceAllString(name, " ")
	return name
}
