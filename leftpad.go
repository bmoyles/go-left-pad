package leftpad

import (
	"strings"
	"unicode/utf8"
)

func LeftPad(str string, length int, chr rune) string {
	// zero-value of an empty rune
	if chr == 0 {
		return str
	}
	strlen := utf8.RuneCountInString(str)
	if strlen >= length {
		return str
	}
	count := length - strlen
	pad := strings.Repeat(string(chr), count)
	return pad + str
}
