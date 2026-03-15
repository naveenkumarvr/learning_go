package main

import "fmt"

func main() {

	var a int = 12
	var b int = 25

	fmt.Printf("Bitwise AND takes two input perform AND operation: %v \n", a&b)
	fmt.Printf("Bitwise OR takes two input perform AND operation: %v \n", a|b)
	// The XOR is 1 when two binary number are different
	fmt.Printf("Bitwise XOR takes two input perform AND operation: %v \n", a^b)
	fmt.Printf("Bitwise Left shift takes Shift binary number 1 place left: %v \n", a<<1)
	fmt.Printf("Bitwise Right shift takes Shift binary number 1 place right: %v \n", a>>1)
}
