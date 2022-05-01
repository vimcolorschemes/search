package str

import (
	"regexp"
	"strings"
)

// Normalize removes most unsafe characters to use in a regular expression
func Normalize(value string) string {
	remove := regexp.MustCompile(`[\{\}\[\]\(\)\\\$\^]`)
	value = remove.ReplaceAllString(value, "")

	replace := regexp.MustCompile(`/`)
	value = replace.ReplaceAllString(value, " ")

	escape := regexp.MustCompile(`([\+\|\?])`)
	value = escape.ReplaceAllString(value, `\$1`)

	return strings.Trim(value, " ")
}
