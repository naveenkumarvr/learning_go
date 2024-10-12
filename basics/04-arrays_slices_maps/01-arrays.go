// It is a collection of similar data elements stored in a contiguous memory location
// we cannot store different data type in single array. Which is called Homogeneous
// They are fixed length. Once declared with the length we cannot change

package main

import "fmt"

func main() {

	// var <array name> <size of array> <data type>
	var grades [5]int = [5]int{1, 2, 3, 4, 5}

	// Array value should be less than or equal to array length not greater than length
	var fruits [3]string = [3]string{"Grapes", "Apple"}

	fmt.Println(grades)
	fmt.Println(fruits)

	//short declaration of array
	marks := [3]int{1, 2, 3}
	fmt.Println(marks)

	//Initialization using elipses [...]
	// By this way we don't want to declare the length of the array
	marks_1 := [...]int{4, 5}
	fmt.Println(marks_1)

	//Properties
	//len()
	fmt.Println(len(marks_1))

	//access array element
	fmt.Println(marks[2])

	//Change array value
	marks[2] = 8
	fmt.Println(marks[2])

	//Looping through an arra
	for i := 0; i < len(marks); i++ {
		fmt.Printf("The value on index %v is %v \n", i, marks[i])
	}

	//Forloop with range
	for index, element := range marks {
		fmt.Println(index, "=>", element)
	}

	// Multi Dimensional arrays
	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr[1][1])

}
