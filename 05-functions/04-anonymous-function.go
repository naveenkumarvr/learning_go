// Anonymous function is a function that is declared without any named identifier to refer to it
// They can accept input and return output
// where it is used ? They can be used for containing functionality that need not be named and possibly for short term use

package main

import "fmt"

func main() {
	x := func(l int, b int) int {
		return l * b
	}
	fmt.Printf("The datatype of the function x is %T \n", x)
	fmt.Printf("Result of 20 * 20 is : %v \n", x(20, 20))

	// Here we are not calling the function separately. Directly we are passing the arg at the end of the function
	y := func(l int, b int) int {
		return l * b
	}(20, 20)
	fmt.Printf("The datatype of the function y is %T \n", y)
	fmt.Printf("Result of 20 * 20 is : %v \n", y)

}
