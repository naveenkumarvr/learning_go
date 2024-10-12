// Interfaces

/*
An interface specifies a method set and is a powerful way to introduce modularity in Go.

Interface is like a blueprint for a method set

They describe all the method of a method set by providing the function signature for each method

They specify a set of methods but do not implement them.

Syntax:
type <interface name> interface{
// method signatures
}

Implementing interfaces
A type implements an interface by implementing its methods. For example structure type.
Go language interfaces are implemented implicitly , it does not have any specific keyword to implement interface
*/

package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

// As per the definition you can see that the struct with below function can be called as shp interface type as it contains the all methods listed in interface
type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perimeter() float64 {
	return 4 * s.side
}

// As per the definition you can see that the struct with below function can be called as interface type as it contains the all methods listed in interface
type rect struct {
	length, breath float64
}

func (r rect) area() float64 {
	return r.length * r.breath
}
func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.breath
}

func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	r := rect{length: 3, breath: 4}
	c := square{side: 5}
	printData(r)
	printData(c)
}
