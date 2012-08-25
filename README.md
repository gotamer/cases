Go String Manipulation
=======

#Uc
Uc returns a copy of the string s with the first unicode words first letter mapped to it's upper case  

#UcAll
UcAll returns a copy of the string s with all unicode words first letter mapped to their upper case


```go
package main

import "gotamer/strings"

func main(){
	word := "uc returns a copy of the string s with the first unicode words first letter mapped to it's upper case\n"
	word = strings.Uc(word)
	print(word)

	satz := "ucAll returns a copy of the string s with all unicode words first letter mapped to their upper case\n"
	satz = strings.UcAll(satz)
	print(satz)
}
```