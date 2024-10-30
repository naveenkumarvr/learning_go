/*
Main function in the go routine is the main go routine and it calls other go routine in other program
Once main go routine exits that means the entire program exited
Go routine doesn't have parents or children. when you execute all executes at the same time
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	go start()
	time.Sleep(1 * time.Second)
}

func start() {
	go process()
	fmt.Println("In Start")
}
func process() {
	fmt.Println("In Process")
}

// SOmetimes Process will print before start which shows that go routine are independent
