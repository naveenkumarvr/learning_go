/* Type Casting: The process of converting one datatype to another is called type casting
 */

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// int to float
	var i int = 20
	var f float64 = float64(i)
	fmt.Printf("the number %v is of type %v is converted to %v", i, reflect.TypeOf(i), reflect.TypeOf(f))

	// Float to integers
	var j float64 = 22.33
	var k int = int(j)
	fmt.Printf("\n%v\n", k)

	//String conversion package strconv
	//	Itoa() >> Convert int to string

	var a int = 45
	var b string = strconv.Itoa(a)
	fmt.Printf("%q", b)

	//Atoi Convert string to int and returns two value output and error
	var c string = "200"
	d, err := strconv.Atoi(c)
	fmt.Printf("%v, %T", d, d)
	fmt.Printf("%v, %T", err, err)

}
