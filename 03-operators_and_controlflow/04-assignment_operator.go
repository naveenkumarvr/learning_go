package main

import "fmt"

func main() {

	var a int = 22
	a += 10
	fmt.Printf("This is assign and add: %v \n", a)
	a -= 10
	fmt.Printf("This is assign and Sub: %v \n", a)
	a *= 2
	fmt.Printf("This is assign and Mul: %v \n", a)
	a /= 2
	fmt.Printf("This is assign and Div and map quotient: %v \n", a)
	a %= 3
	fmt.Printf("This is assign and Modulus: %v \n", a)
}
