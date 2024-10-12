package main

import "fmt"

func main() {

	var city string = "Chennai"
	var user string = "Naveen"
	var num int = 42

	fmt.Print("Welcome to ", city, ", ", user)

	//string formatting
	/* Format Specifier
	- printf >> Format Specifier
	- The format specifier keyword always comes with '%'
		- %v >> formats the value in a default format
		- %d >> Format decimal integers
		- %T >> Type of value
		- %c >> Character
		- %q >> Quoted value
		- %s >> Plain string
		- %t >> ture or false
		- %f >> Floating numbers
		- %.2f >> Floating number with 2 decimal point
	*/
	fmt.Printf("Template string %v", user)
	fmt.Printf("Makrs: %d", num)
	fmt.Printf("Using it together %v, and your number is %d", user, num)
}

// \n newline charcterc
