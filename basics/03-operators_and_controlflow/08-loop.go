package main

import "fmt"

func main() {

	for i := 1; i <= 3; i++ {
		fmt.Println("Hello World")
	}

	j := 1

	for j <= 5 {
		fmt.Println(j * j)
		j++
	}

	// Break and Continue
	for k := 1; k <= 5; k++ {
		if k == 3 {
			break
		}
		fmt.Println(k)
	}

	for k := 1; k <= 5; k++ {
		if k == 3 {
			continue
		}
		fmt.Println(k)
	}

}
