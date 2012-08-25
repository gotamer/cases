// uc.go
package strings

import "strings"

func Uc(s string) (uc string) {
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
