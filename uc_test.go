package strings

import "testing"

func TestWord(t *testing.T) {
	satz := "ucWord returns a copy of the string s with the first unicode word mapped to it's upper case\n"
	satz = Uc(satz)
	print(satz)
}

func TestWords(t *testing.T) {
	satz := "ucWords returns a copy of the string s with all unicode words mapped to their upper case\n"
	satz = UcAll(satz)
	print(satz)
}
