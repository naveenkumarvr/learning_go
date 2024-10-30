// Passing by value in function
// Function is called by directly passing the value of the variable as an argument. i.e the parameter is copied into another location of the memory
// So when accessing or modifying the variable within your function, only the copy is accessed or modified and the original value never changed.
// All Basic types (int, float, bool, string, array) are passed by value

package main

import "fmt"

func modify(s string) {
	s = "world"

}

func main() {

	a := "hello"
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
}
