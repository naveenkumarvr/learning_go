/*
Select the correct base case for function printSquares -
Select the correct base case for function printSquares -


for input n -> prints squares for n, n-1, n-2, …… -5



Example:



Input: n=2; Output: 4 1 0 1 4 9 16 25

package main

import "fmt"

func printSquares(n int) {
    // base case

    fmt.Printf("%d ", n*n)
    printSquares(n - 1)
}

func main() {
    printSquares(2)
}
*/

package main

import "fmt"

func printSquares(n int) {

	if n < -5 {
		return
	}

	fmt.Printf("%d ", n*n)
	printSquares(n - 1)

}

func main() {
	printSquares(2)
}
