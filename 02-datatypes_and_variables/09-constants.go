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

}
