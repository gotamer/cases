// Provides conversion and case manipulation, 
// such as upper case word, upper case words, 
// camel case, dash, underscore etc.  
package cases

import "strings"
import "unicode"

var is bool

// Uc returns a copy of the string s with the first unicode words, 
// first letter mapped to it's upper case  
func Uc(s string) string {
	s = strings.TrimRight(s, " -_\t")
	s = strings.TrimLeft(s, " -_\t\n\r")
	return strings.ToUpper(s[0:1]) + s[1:]
}

// UcAll returns a copy of the string s with all unicode words,
// first letters mapped to their upper case
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

// Camel returns a copy of the string s with all unicode words,
// first letters mapped to their upper case, and spaces removed 
func Camel(s string) (uc string) {
	for _, word := range strings.Split(s, " ") {
		if word != "" && word != " " && word != "_" && word != "-" && word != "\t" && word != "\n" {
			uc += Uc(word)
		} else {
			uc += ""
		}
	}
	return
}

// UnCamel returns a copy of the string s with all upper case letters seperated by spaces
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

// UnCamel returns a copy of the string s with all upper case letters seperated by underscore
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

// UnderScore returns a copy of the string s with all the words seperated by underscore
func UnderScore(text string) (newtext string) {

	text = strings.Trim(text, " -_\t\n\r")
	is = false
	for _, r := range text {

		if unicode.IsSpace(r) {
			if is != true {
				is = true
				newtext += "_"
			}
		} else {
			is = false
			newtext += string(r)
		}
	}
	return newtext
}

// Dash returns a copy of the string s with all the words seperated by dashes
func Dash(text string) (newtext string) {
	text = strings.Trim(text, " -_\t\n\r")
	is = false
	for _, r := range text {
		if unicode.IsSpace(r) {
			if is != true {
				is = true
				newtext += "-"
			}
		} else {
			is = false
			newtext += string(r)
		}
	}
	return newtext
}

// UnDash replaces dashes with spaces
func UnDash(t string) (n string) {
	t = strings.Trim(t, " -_\t\n\r")
	for _, r := range t {
		s := string(r)
		if s == "-" {
			s = " "
		}
		n += s
	}
	return
}

// DashUnder returns a copy of the string s with words seperated by dashes 
// or underscore depending on if the word is an upper case or lower case word
func DashUnder(text string) (newtext string) {

	text = strings.Trim(text, " -_\t\n\r")

	for c, r := range text {
		s := string(r)
		if unicode.IsSpace(r) {
			is = true
		} else {
			if is == true || c == 0 {
				if unicode.IsUpper(r) {
					newtext += "_" + s
				} else {
					newtext += "-" + s
				}
			} else {
				newtext += s
			}
			is = false
		}
	}
	return strings.ToLower(newtext)
}

// Du is like DashUnder but it assumes the first letter of the text to be always upper case
func Du(text string) (newtext string) {

	text = strings.Trim(text, " -_\t\n\r")

	for c, r := range text {
		s := string(r)
		if unicode.IsSpace(r) {
			is = true
		} else {
			if is == true {
				if unicode.IsUpper(r) {
					newtext += "_" + s
				} else {
					newtext += "-" + s
				}
			} else {
				newtext += s
			}
			is = false
		}
	}
	return strings.ToLower(newtext)
}

// Udu undoes Du
func Udu(text string) (newtext string) {

	text = strings.Trim(text, " -_\t\n\r")
	is = false
	for c, r := range text {
		s := string(r)
		if c == 0 {
			newtext += strings.ToUpper(s)
		} else if s == "-" {
			is = false
			newtext += " "
		} else if s == "_" {
			is = true
			newtext += " "
		} else {
			if is == true {
				newtext += strings.ToUpper(s)
			} else {
				newtext += s
			}
			is = false
		}
	}
	return strings.Trim(newtext, " -_")
}

// UnDashUnder undoes DashUnder
func UnDashUnder(text string) (newtext string) {

	text = strings.Trim(text, " -_\t\n\r")
	is = false
	for _, r := range text {
		s := string(r)
		if s == "-" {
			is = false
			newtext += " "
		} else if s == "_" {
			is = true
			newtext += " "
		} else {
			if is == true {
				newtext += strings.ToUpper(s)
			} else {
				newtext += s
			}
			is = false
		}
	}
	return strings.Trim(newtext, " ")
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
