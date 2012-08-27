// cases.go
package cases

import "strings"
import "unicode"

var is bool

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
	is = false
	for _, word := range strings.Split(s, " ") {
		if word != "" && word != " " && word != "\t" && word != "\n" {
			is = false
			uc += " " + Uc(word)
		} else {
			is = true
		}
	}
	uc = strings.Trim(uc, " -_")
	return
}

func Camel(s string) (uc string) {
	for _, word := range strings.Split(s, " ") {
		if word != "" && word != " " && word != "_" && word != "-" {
			uc += Uc(word)
		} else {
			uc += ""
		}
	}
	return
}

func UnCamel(text string) (newtext string) {
	is = false
	for _, s := range text {
		if isSeparator(s) {
			if is != true {
				newtext += " " + string(s)
			} else {
				continue
			}
		} else {
			newtext += string(s)
		}
	}
	return newtext
}

func UnCamelUnderScore(text string) (newtext string) {
	is = false
	for _, s := range text {
		if isSeparator(s) {
			if is != true {
				newtext += "_" + string(s)
			} else {
				continue
			}
		} else {
			is = false
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
