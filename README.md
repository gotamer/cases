Go String Case Manipulation
===========================

###`cases.Uc(s string)`
Uc returns a copy of the string s with the first unicode words, first letter mapped to it's upper case  

###`cases.UcAll(s string)`
UcAll returns a copy of the string s with all unicode words, first letters mapped to their upper case

### CamelCase `cases.Camel(s string)` 
Camel returns a copy of the string s with all unicode words, first letters mapped to their upper case, and spaces removed 

### Un Camel Case `cases.UnCamel(s string)` 
UnCamel returns a copy of the string s with all upper case letters seperated by spaces
