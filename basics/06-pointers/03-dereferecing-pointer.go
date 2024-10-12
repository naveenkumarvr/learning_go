/*
Syntax : *<Name of the pointer>
*/

package main

import "fmt"

func main() {

	s := "Hello"
	fmt.Printf("%T, %v \n", s, s)

	ps := &s
	*ps = "world" // Here we are dereferencing the pointer and changing its value which changes the variable value as well
	fmt.Printf("%T, %v \n", s, s)
}
