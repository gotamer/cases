// cases.go
package cases

import "strings"
import "unicode"

func Uc(s string) (uc string) {
	s = strings.Trim(s, " -_\t")
	s = strings.ToLower(s)
	for pos, char := range strings.Split(s, "") {
		if pos == 0 {
			uc += strings.ToUpper(char)
			continue
		}
		uc += char
	}
	return
}

func UcAll(s string) (uc string) {
	for _, word := range strings.Split(s, " ") {
		if word != "" && word != " " {
			uc += Uc(word)
		}
		uc += " "
	}
	return
}

func Camel(s string) (uc string) {
	for _, word := range strings.Split(s, " ") {
		if word != "" && word != " " && word != "_" && word != "-" {
			uc += Uc(word)
		}
		uc += ""
	}
	return
}

func UnCamel(text string) (newtext string) {
	for _, s := range text {
		if isSeparator(s) {
			newtext += " " + string(s)
		} else {
			newtext += string(s)
		}
	}
	return newtext
}

// isSeparator reports whether the rune could mark a word boundary.
// TODO: update when package unicode captures more of the properties.
func isSeparator(r rune) bool {
	if r <= 0x7F {
		// Digits are separators
		if unicode.IsDigit(r) {
			return true
		}
		switch {
		case '0' <= r && r <= '9':
			return true
		case 'A' <= r && r <= 'Z':
			return true
		}
	}
	return false
}
