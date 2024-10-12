/*
Struct is a user defined data type
It is  a structure that groups together data elements
Provide a way to reference a series of grouped values through single variable name
used when it makes sense to group or associate two or more data variable

SYNTAX
Declaration
type <struct_name> struct {
	// list of fields
}

Initialization
var <variable name> <struct name>
//Alternative initialization
<variable name> := new(<structname>)
//Alternative initialization
<variable name> := <struct name> {
	<field_name>: >value>,
	<field_name>: >value>,
}

*/

package main

import "fmt"

// Declaration
type Student struct {
	name   string
	rollNo int
}

func main() {

	//Initialization
	var s Student
	fmt.Printf("%+v \n", s)

	//Alternative initialization
	st := new(Student)
	fmt.Printf("%+v \n", st)

	//Alternative initialization
	st_1 := Student{
		name:   "Joe",
		rollNo: 25,
	}
	fmt.Printf("%+v \n", st_1)

	//Alternative initialization
	st_2 := Student{"Loe", 25}
	fmt.Printf("%+v \n", st_2)

}
