// Accessing fields of the struct

package main

import "fmt"

type Circle struct {
	x      int
	y      int
	radius int
}

func main() {
	var c Circle
	c.x = 5
	c.y = 5
	c.radius = 5

	fmt.Printf("%+v", c)
}
