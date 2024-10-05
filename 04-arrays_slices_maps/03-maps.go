// MAPS are unordered collection of Key value pair
// They are implemented by Hash Table

package main

import "fmt"

func main() {
	// Syntax: var <mapname> map[<data type of key>]<value data type>
	// The zero value of the map is nil and it doesn't contain any key
	// we cannot add a value to a Nil Map
	var codes map[string]string
	fmt.Println(codes)

	// Map with initialization
	// Syntax:  <mapname>:= map[<data type of key>]<value data type>
	codes_1 := map[string]string{"en": "English", "fr": "French"}
	fmt.Println(codes_1)

	// Declaring with make function
	// Syntax: mapname := make(map[<data type of key>]<data type of value>, <Initial capacity>)
	// Note capacity is optional
	codes_2 := make(map[string]int)
	fmt.Println(codes_2)

	// Determine the length of map
	fmt.Println(len(codes_1))

	// Get value for the key
	fmt.Println(codes_1["en"])

	// Getting a key validation.
	// By default the map return two value 1. value 2. Boolean which says whether the key is present on not
	val, found := codes_1["fr"]
	fmt.Println(val, "=", found)
	val1, found1 := codes_1["abc"]
	fmt.Println(val1, "=", found1)

	// Adding new entry to a map
	codes_1["it"] = "Italian"
	fmt.Println(codes_1)

	// Update the item
	codes_1["en"] = "British English"
	fmt.Println(codes_1)

	// Delete the keyvalue pair
	delete(codes_1, "en")
	fmt.Println(codes_1)

	//Looping, the range will return key and value
	for key, value := range codes_1 {
		fmt.Println(key, "=", value)
	}

	// Truncating a map
	codes_1 = make(map[string]string)
	fmt.Println(codes_1)

}
