package main

import (
	"fmt"
	"reflect"
)

func main() {

	var grade int = 32
	var isCheck bool = true
	var amount float32 = 33.33
	var message string = "hello Naveen"

	// Using %T
	fmt.Printf("The Grade is = %v is of type %T \n", grade, grade)
	fmt.Printf("The Grade is = %v is of type %T \n", message, message)
	fmt.Printf("The Grade is = %v is of type %T \n", amount, amount)
	fmt.Printf("The Grade is = %v is of type %T \n", isCheck, isCheck)

	// using reflect.Typeof()
	fmt.Printf("The Grade is = %v \n", reflect.TypeOf(1000))
	fmt.Printf("The Grade is = %v  \n", reflect.TypeOf("Naveen"))
	fmt.Printf("The Grade is = %v  \n", reflect.TypeOf(46.0))
	fmt.Printf("The Grade is = %v  \n", reflect.TypeOf(true))
	//Typeof with variable
	fmt.Printf("The Grade is = %v is of type %v \n", amount, reflect.TypeOf(amount))
}
