// Passing struct to function

package main

import "fmt"

type Circle struct {
	x      int
	y      int
	radius float64
	area   float64
}

func calcArea(c *Circle) {
	const PI float64 = 3.14
	var area float64
	area = (PI * c.radius * c.radius)
	//c.area = area
	(*c).area = area
}

func main() {
	c := Circle{x: 5, y: 5, radius: 5, area: 0}
	fmt.Printf("%+v \n", c)
	calcArea(&c)
	fmt.Printf("%+v \n", c)
}
