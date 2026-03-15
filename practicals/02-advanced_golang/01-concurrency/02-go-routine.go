/*
Example
*/

package main

import (
	"fmt"
	"time"
)

func calculateSqure(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
}

func main() {
	start := time.Now()
	for i := 1; i <= 10000; i++ {
		go calculateSqure(i)
	}
	elapsed := time.Since(start)
	time.Sleep(2 * time.Second) // This will alow program to wait and show the results
	fmt.Println("Function Took : ", elapsed)
}
