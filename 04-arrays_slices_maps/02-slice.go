/*
Continuous segment of an underlying array
Unlike array slices are variable typed (we don't want to declare size up front)
*/

package main

import "fmt"

func main() {
	// Slice should have pointer, length and capacity

	// Syntax:
	// <slicename> := []<data_type>{<value>}
	grades := []int{10, 20, 30}
	fmt.Println(grades)

	// create slice from an arrya
	// array[start_index : end_index]
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slice_1 := arr[1:5]
	fmt.Println(slice_1)

	//Subslice
	sub_slice := slice_1[1:2]
	fmt.Println(sub_slice)

	// Declaring slice via make function. Mostly used to create empty slice
	// Syntax:
	// <slice name> := make([]<data_type>, length, capacity)
	slice_make := make([]int, 3, 4)
	fmt.Println(len(slice_make))
	fmt.Println(cap(slice_make))
	fmt.Println(slice_make)

	// Slice index numbers
	// When you change the value of a slice the underlying array will also get affected.
	// Slice is just a reference from an array
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice_2 := arr1[:3]
	fmt.Println(slice_2)
	slice_2[2] = 20
	fmt.Println(arr1)
	fmt.Println(slice_2)

	// Appending new element to slice
	// Syntax <slicename> = append(slice, element1, element2)
	fmt.Println(len(slice_2))
	fmt.Println(cap(slice_2))
	slice_2 = append(slice_2, 20, 30, 40, 59)
	fmt.Println(slice_2)
	fmt.Println(len(slice_2))
	fmt.Println(cap(slice_2))
	//Appending slice to another slice
	slice_2 = append(slice_2, slice_1...)
	fmt.Println(slice_2)

	//delete element from slice
	slice_1 = arr1[:2]
	slice_2 = arr1[3:5]
	slice_4 := append(slice_1, slice_2...)
	fmt.Println(slice_4)

	// Copy one slice to another
	// syntax : x = copy(dest_slice, src_slice)
	// Both slice should be of same datatype
	slice_5 := copy(slice_4, slice_2)
	fmt.Println(slice_5) // This will not print the slice it prints the number of variable in the slice

	//Looping through slice is same as looping through array
	for index, element := range slice_4 {
		fmt.Println(index, "=>", element)
	}
	// When only needed the value in for loop
	for _, value := range slice_4 {
		fmt.Println(value)
	}

}
