/*
Complete the code in the function calculate to return a slice consisting of 4 elements [ sum of a and b, difference of a and b, product of a and b, quotient on dividing a by b]
A Go file is located at /root/code/calculator/ directory.
package main

import "fmt"

	func calculate(a int, b int) []float64 {
	    // your code goes here
	    return results
	}

	func main() {
	    fmt.Println(calculate(20, 10))
	    fmt.Println(calculate(700, 70))
	}
*/
package main

import "fmt"

func calculate(a int, b int) []float64 {
	results := []float64{float64(a) + float64(b), float64(a) - float64(b), float64(a) * float64(b), float64(a) / float64(b)}
	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}
