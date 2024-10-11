/*
Go support pointers allowing you to pass references to values within your program

In call by reference/pointers the address of the variable is passed into the function call as the actual parameter

All the operation in the function are performed on the value stored at the address of the actual parameter
*/

package main

import "fmt"

func modify(s *string) {
	*s = "world"
}

func modify_1(s []int) {
	s[0] = 100
}

func modify_ascii(m map[string]int) {
	m["K"] = 75
}

func main() {
	a := "hello"
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)

	// Slices and maps are passed by reference by default
	slice := []int{10, 20, 30}
	fmt.Println(slice)
	modify_1(slice)
	fmt.Println(slice)

	// Example with maps
	ascii_codes := make(map[string]int)
	ascii_codes["A"] = 65
	ascii_codes["B"] = 70
	fmt.Println(ascii_codes)
	modify_ascii(ascii_codes)
	fmt.Println(ascii_codes)
}
