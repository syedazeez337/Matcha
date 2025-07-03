package tokenize

import (
	"regexp"
	"strings"
)

func Tokenize(text string) []string {
	// Remove punctuation
	re := regexp.MustCompile(`[^\w]+`)
	cleaned := re.ReplaceAllString(text, " ")

	// convert to lowercase
	cleaned = strings.ToLower(cleaned)

	// split into tokens(words)
	tokens := strings.Fields(cleaned)

	return tokens
}