/*
We are working on a project for a grocery store. The store has a list of items that are currently in stock, and we need to implement a function that will return a list of items that are in a given price range.

package main

import "fmt"

type Item struct {
    Name  string
    Price float64
}

func getItemsInPriceRange(items []Item, minPrice, maxPrice float64) []Item {
    // your code goes here
}

func main() {
    items := []Item{
        {Name: "Apple", Price: 0.5},
        {Name: "Banana", Price: 0.25},
        {Name: "Orange", Price: 0.75},
        {Name: "Pineapple", Price: 1.5},
    }

    fmt.Println(getItemsInPriceRange(items, 0.0, 1.0))
    fmt.Println(getItemsInPriceRange(items, 0.5, 1.0))
    fmt.Println(getItemsInPriceRange(items, 0.75, 1.5))
}

*/

package main

import (
	"fmt"
)

type Item struct {
	Name  string
	Price float64
}

func getItemsInPriceRange(items []Item, minPrice, maxPrice float64) []Item {
	// your code goes here
	var item []Item
	for _, value := range items {
		if (value.Price >= minPrice) && (value.Price <= maxPrice) {
			item = append(item, value)
		}

	}
	return item
}

func main() {
	items := []Item{
		{Name: "Apple", Price: 0.5},
		{Name: "Banana", Price: 0.25},
		{Name: "Orange", Price: 0.75},
		{Name: "Pineapple", Price: 1.5},
	}

	fmt.Println(getItemsInPriceRange(items, 0.0, 1.0))
	fmt.Println(getItemsInPriceRange(items, 0.5, 1.0))
	fmt.Println(getItemsInPriceRange(items, 0.75, 1.5))
}
