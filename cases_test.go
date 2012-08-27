package cases

import "testing"

func TestWord(t *testing.T) {
	print("\n")
	satz := "case.Uc returns a copy of the string s with the first unicode word mapped to it's upper case\n"
	satz = Uc(satz)
	print(satz)
}

func TestWords(t *testing.T) {
	print("\n")
	satz := "case.UcAll returns a copy of the string s with all unicode words mapped to their upper case\n"
	satz = UcAll(satz)
	print(satz)
}

func TestCamel(t *testing.T) {
	print("\n")
	satz := "case.Camel returns a copy of the string s with all unicode words mapped to their upper case, and sapaces removed\n"
	satz = Camel(satz)
	print(satz)
}
