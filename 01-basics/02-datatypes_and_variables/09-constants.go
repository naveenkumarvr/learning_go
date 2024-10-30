package main

import "fmt"

func main() {
	// Constant values will not be changed once declared
	// Syntax: const <const name> <data type> = <value>, datatype is optional it is not necessary to declare the datatype
	// untyped constant >> where no datatype is declared, typed constant >> where data type is declared

	// Untyped constant
	const age = 12
	const name, num = "naveen", 32

	// Typed constant
	const myName string = "my Name"

	fmt.Printf("%v", age)
	fmt.Printf("%v", name)
	fmt.Printf("%v", myName)

	/*
		Some key factory
		- we are not allowed to declare a constant without value
		- Short form declaration using := is not allowed with const
		- Once declared the values cannot be changed
	*/

}
