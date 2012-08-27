Go String Case Manipulation
=======

###`cases.Uc(s string)`
Uc returns a copy of the string s with the first unicode words, first letter mapped to it's upper case  

###`cases.UcAll(s string)`
UcAll returns a copy of the string s with all unicode words first letter mapped to their upper case

### CamelCase `cases.Camel(s string)` 
Camel returns a copy of the string s with all unicode words first letter mapped to their upper case, and saces removed 


###`cases.Uc(s string)`
Uc returns a copy of the string s with the first unicode words, first letter mapped to it's upper case  

```go
package main

import "gotamer/cases"

func main(){
	word := "uc returns a copy of the string s with the first unicode words, first letter mapped to it's upper case\n"
	word = strings.Uc(word)
	print(word)

}
```
###Result:  
Run start ...

Uc returns a copy of the string s with the first unicode words first letter mapped to it's upper case  
 
Run finished 0, process exited normally.


###`cases.UcAll(s string)`
UcAll returns a copy of the string s with all unicode words first letter mapped to their upper case


```go
package main

import "gotamer/strings"

func main(){
	satz := "ucAll returns a copy of the string s with all unicode words first letter mapped to their upper case\n"
	satz = strings.UcAll(satz)
	print(satz)
}
```
###Result:  
Run start ...

UcAll Returns A Copy Of The String S With All Unicode Words First Letter Mapped To Their Upper Case  
 
Run finished 0, process exited normally.
