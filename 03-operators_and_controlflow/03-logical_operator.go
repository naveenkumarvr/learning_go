package main

import "fmt"

func main() {

	var a int = 10

	fmt.Printf("Logical And operator &&:  %v \n", (a < 100) && (a < 200))
	fmt.Printf("Logical OR operator ||: %v \n", (a < 100) || (a < 9))
	fmt.Printf("Logical NOT Operator !: %v \n", !(a > 15))

}
