// Function Syntax
/*
func <function name>(<params>)<return type>{
	// Function body
}
*/

package main

import "fmt"

func addNumbers(a int, b int) int {
	sum := a + b
	return sum

}

func main() {
	// Calling a function
	// Syntax <function name>(<arguments>)
	sumOfNumbers := addNumbers(2, 3)
	fmt.Println(sumOfNumbers)
}
