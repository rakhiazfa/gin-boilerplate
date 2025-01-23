package utils

import (
	"regexp"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func LcFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func ToTitleCase(s string) string {
	regex := regexp.MustCompile(`([a-z])([A-Z])`)
	withSpaces := regex.ReplaceAllString(s, `$1 $2`)

	titleCaser := cases.Title(language.English)
	return titleCaser.String(withSpaces)
}
