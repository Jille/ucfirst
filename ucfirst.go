package ucfirst

import (
	"unicode"
	"unicode/utf8"
)

// Ucfirst converts the first rune to uppercase and leaves the rest unchanged.
func Ucfirst(s string) string {
	r, l := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		// Empty string or invalid UTF8
		return s
	}
	r = unicode.ToTitle(r)
	return string(r) + s[l:]
}
