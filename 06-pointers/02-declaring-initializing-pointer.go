// Syntax: var <pointer name> *<data type>

package main

import "fmt"

func main() {
	var ptr_i *int
	var ptr_s *string

	fmt.Println(ptr_i)
	fmt.Println(ptr_s)

	//Initializing the pointer
	// Syntax : var <pointer name> *<data type> = &<variable name>
	i := 10

	var ptr_ii *int = &i
	fmt.Println(ptr_ii)

	// Alternative method to declare pointer
	// synatx: var <pointer name> = &<variable name>
	s := "hello"
	var ptr_ss = &s
	fmt.Println(ptr_ss)

	// Alternative method using shorthand
	//syntax: <poniter_name> := &<varaible_name>
	q := "SayHello"
	ptr_q := &q
	fmt.Println(ptr_q)

}
