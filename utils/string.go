package utils

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

func Capitalize(s string) string {
	if s == "" {
		return s
	}

	r, size := utf8.DecodeRuneInString(s)

	if !unicode.IsLower(r) {
		return s
	}

	buf := &bytes.Buffer{}
	buf.WriteRune(unicode.ToUpper(r))
	buf.WriteString(s[size:])

	return buf.String()
}
