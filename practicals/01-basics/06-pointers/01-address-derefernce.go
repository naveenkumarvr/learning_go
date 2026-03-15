/*
Pointers are variable that holds the memory address of other variable
They point where the memory is allocated and provide ways to find or even change the value located at the memory location

Address: & :The address of the variable can be obtained by preceding the name of the variable with an Ampersand sign (&). This is known as Address-of-operator
Deference: * : It is known as dereference operator. Whe placed before address it returns the value at that address

*/

package main

import "fmt"

func main() {
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

}
