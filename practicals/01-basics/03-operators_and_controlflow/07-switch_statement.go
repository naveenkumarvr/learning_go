package main

import "fmt"

func main() {

	var i int = 100

	switch i {
	case 10:
		fmt.Println("i is 10")
	case 100, 200:
		fmt.Println("i is 100 or 200")
		fallthrough // this keyword is used in switch-case to force the execution flow to fall through the successive case block
	case 300:
		fmt.Println("i is 300")
		fallthrough // this keyword is used in switch-case to force the execution flow to fall through the successive case block
	default:
		fmt.Println("i is neither 10 or 100, 200")
	}

	var a, b int = 10, 20
	switch {
	case a+b == 30:
		fmt.Println("A is 30")
	case a+b <= 30:
		fmt.Println("A less than 30")
	default:
		fmt.Println("A is greater than 30")

	}

}
