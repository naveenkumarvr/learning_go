// Functions calls itself by direct or indirect means
// functions keeps calling itself until it reaches a base case
// Where it is used ? Generally used to solve a problem where the solution is dependent on the smaller instance of the same problem

package main

import "fmt"

// Example factorial of number eg. fact(3) 3*2*1

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	n := 5
	result := factorial(n)
	fmt.Println(result)
}
