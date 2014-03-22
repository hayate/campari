package helpers

import (
	"strings"
	"unicode"
)

func Capitalize(s string) string {
	s = strings.ToLower(s)
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
