package main

import "fmt"


var name string = "Global Variable" // This is a global variable which can be accessed all over the program

func main() {
	// inner blocks can access variable from outer block but outer block cannot access variable from inner block
	city := "Chennai" // This is also called as Local variable. Local to the function
		country := "India"
		fmt.Println(city)
		fmt.Println(country)
	}
	fmt.Println(city)
	// fmt.Println(country)


}
