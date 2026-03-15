package main

import "fmt"

func main() {

	// var name string
	// fmt.Print("Enter the name: ")
	// fmt.Scanf("%s", &name)
	// fmt.Println(name)

	// // Multiple inputs from scanf
	// var name_1 string
	// var is_muggle bool

	// fmt.Print("Enter your name and are you muggle: ")
	// fmt.Scanf("%s %t", &name_1, &is_muggle)
	// fmt.Println(name_1, is_muggle)

	//fmt.Scan returns two value, Count>>Number of arg that function write to and Error>> Error during the execution
	var a string
	var b int

	fmt.Print("Enter the string and number: ")
	count, err := fmt.Scanf("%s, %d", &a, &b)
	fmt.Println("Count: ", count)
	fmt.Println("error: ", err)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
