/*
We are building a program to keep track of the number of visitors to a website.


We need to store the number of active visitors in a variable and update it each time a new visitor arrives or an old visitor leaves the website.


A Go file is located at /root/code/visitor directory.

package main

import "fmt"

// Declare variable activeUserCount
// your code goes here

func entry() {
    // Hint: you can use the "++" operator to increment a variable by 1
    // your code goes here
}

func exit() {
    // Hint: you can use the "--" operator to decrement a variable by 1
    // your code goes here
}

func main() {
    entry()
    entry()
    exit()
    exit()
    entry()
    entry()
    fmt.Println(activeUserCount)
}
*/

package main

import "fmt"

// Declare variable activeUserCount
// your code goes here

func entry(count int) int {
	// Hint: you can use the "++" operator to increment a variable by 1
	// your code goes here
	count++
	return count
}

func exit(count int) int {
	// Hint: you can use the "--" operator to decrement a variable by 1
	// your code goes here
	count--
	return count
}

func main() {
	var activeUserCount int = 0
	activeUserCount = entry(activeUserCount)
	activeUserCount = entry(activeUserCount)
	activeUserCount = exit(activeUserCount)
	activeUserCount = exit(activeUserCount)
	activeUserCount = entry(activeUserCount)
	activeUserCount = entry(activeUserCount)
	fmt.Println(activeUserCount)
}
