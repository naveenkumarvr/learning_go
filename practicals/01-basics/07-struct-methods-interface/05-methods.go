// Methods
/*
A method augments a function by adding an extra parameter section immediately after the 'func' keyword that accepts a single arguments

This argument is called 'Receiver'

When to use it ? >> Whenever you feel that there is a strong relation between your struct and func you can use method

Syntax:
func <receiver name> funcationname(){
// Body
}
*/

package main

import "fmt"

type Circle struct {
	radius float64
	area   float64
}

// Receiver as Pointer
func (c *Circle) calcArea() {
	// Here you can see that the receiver is the pointer to the Struct
	c.area = 3.14 * c.radius * c.radius
}

// Receiver as Instance
func (c Circle) calcArea_I() {
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	c := Circle{radius: 5}
	c.calcArea()
	fmt.Printf("%+v \n", c)

	d := Circle{radius: 5}
	d.calcArea_I()
	fmt.Printf("%+v \n", d) // here the value of the original struct will not change as it just an instance not a pointer
}
