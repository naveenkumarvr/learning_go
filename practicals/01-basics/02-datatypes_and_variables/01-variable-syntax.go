//Kinds and Types
/*
Integers
int
	- uint >> unsinged integers. Only positive integers
	- int >> Signed integers contains both + and negatie
float
	- Float 64 is recommended.
	- 2 Decimal points
	- Float of two tyeps float32 and float 64
String
	- defined by "string". Occupies 16byte of memory

Boolean
	- defined by "bool".
	- true or false
*/

// Variable Syntax Declaration
// syntax >> var <variable name> <data type> = <value>

package main

import "fmt"

func main() {
	var myString string = "Hello"
	var myInt int = 32
	var myBool bool = true
	var myFloat float64 = 32.30

	fmt.Println(myString, myInt, myBool, myFloat)
}
