package main

import "fmt"

func main() {

	var a int = 12
	var b int = 5

	fmt.Println(a + b)

	var c, d string = "Hello", "Naveen"
	fmt.Printf("Addition: %d \n", a+b)
	fmt.Printf("Addition String: %v \n", c+d)
	fmt.Printf("Subtraction: %d \n", a-b)
	fmt.Printf("Multiplication: %v \n", a*b)
	fmt.Printf("Division: %v \n", a/b)
	fmt.Printf("Modulus: %v \n", a%b)
}
