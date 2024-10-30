package main

import "fmt"

func main() {

	var s, t string = "Naveen", "Kumar" // Short hand of same data type
	fmt.Println(s, t)
	// fmt.Println(t)

	var (
		h   string = "Naveen"
		age int    = 32
	) // variables of different datatype can be used inside var
	fmt.Println(h, age)

	// Short Variable Declaration
	g := "Hello World"
	fmt.Println(g)
}
