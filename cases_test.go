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
	satz := "case.UcAll returns a copy of     the string s with all unicode				 words mapped to their upper case\n"
	satz = UcAll(satz)
	print(satz)
}

func TestCamel(t *testing.T) {
	print("\n")
	satz := "case.Camel returns a copy of the string s with all unicode words mapped to their upper case, and sapaces removed\n"
	satz = Camel(satz)
	print(satz)
}

func TestUnCamel(t *testing.T) {
	print("\n")
	satz := "case.UnCamelReturnsACopyOfTheStringSWithTheWordsSeperatedBySpace\n"
	satz = UnCamel(satz)
	print(satz)
}

func TestUnCamelUnderScore(t *testing.T) {
	print("\n")
	satz := "case.UnCamelReturnsACopyOfTheStringSWithTheWordsSeperatedBySpace\n"
	satz = UnCamelUnderScore(satz)
	println(satz)
}

func TestUnderScore(t *testing.T) {
	print("UnderScore Test:\n")
	satz := "case.UnderScore returns a copy   \t of the string \n s with the words seperated by underscore\n"
	println("Going in:    ", satz)
	satz = UnderScore(satz)
	println("Comming out: ", satz)
}

func TestDash(t *testing.T) {
	print("Dash Test:\n")
	satz := "case.Dash returns a copy   \t of the string \n s with the words seperated by dashes\n"
	println("Going in:    ", satz)
	satz = Dash(satz)
	println("Comming out: ", satz)
}

func TestDashUnder(t *testing.T) {
	println("\n\n")
	println("DashUnder Test:")
	satz := "case.Dash returns a copy   \t of the String \n s with the Words seperated by Dashes\n"
	println("Going in:    ", satz)
	satz = DashUnder(satz)
	println("DashUnder: ", satz)
	satz = UnDashUnder(satz)
	println("UnDashUnder: ", satz)
}
