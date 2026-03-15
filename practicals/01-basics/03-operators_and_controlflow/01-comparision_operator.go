package main

import "fmt"

func main() {

	var city string = "Chennai"
	var city_2 string = "Madras"
	fmt.Println(city == city_2)

	fmt.Println(city != city_2)

	var num_1 int = 22
	var num_2 int = 23

	fmt.Println(num_1 < num_2)
	fmt.Println(num_1 <= num_2)
}
