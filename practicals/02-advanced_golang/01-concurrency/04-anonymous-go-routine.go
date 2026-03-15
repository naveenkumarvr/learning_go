/*
Anonymous functions are functions those don't have any name
Simply put anonymous function don't use any variables as name when they are declared
These anonymous function can also be called through go routine

Syntax
go func(){
//body
}
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("In Anonymous function")
	}()
	time.Sleep(1 * time.Second)

}
