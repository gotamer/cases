// cases.go
package cases

import "strings"

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
