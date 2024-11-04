package util

import (
	"regexp"
	"strings"
)

var whitespaceRegex = regexp.MustCompile(`\s+`)

func RemoveExtraWhitespace(text string) string {
	text = strings.TrimSpace(text)
	text = whitespaceRegex.ReplaceAllString(text, " ")

	return text
}

func ReplaceWhitespaces(text string, replacement string) string {
	text = whitespaceRegex.ReplaceAllString(text, replacement)

	return text
}

func FullySanatize(text string) string {
	text = RemoveExtraWhitespace(text)
	text = ReplaceWhitespaces(text, "-")
	text = strings.ToLower(text)

	return text
}
