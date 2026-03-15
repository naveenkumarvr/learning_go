package main

import "fmt"

func main() {
	var a string = "Tryinga"

	if a == "happy" {
		fmt.Println(a)
	} else if a == "Trying" {
		fmt.Println("Trying to be happy")
	} else { // Else should be right next to the if block end
		fmt.Println("Random String")
	}
}
