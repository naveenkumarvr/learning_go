// It delays the execution of the statement until the surrounding function returns
// The defer call args are evaluated immediately but the function call is not executed until the surrounding function returns

package main

import "fmt"

func printName(str string) {
	fmt.Println(str)
}
func printRollNo(rno int) {
	fmt.Println(rno)
}
func printAddress(adr string) {
	fmt.Println(adr)
}

func main() {
	printName("Joe")
	defer printRollNo(25) // This function waits until all the other functions gets executed before running it
	printAddress("Street-01")
}
