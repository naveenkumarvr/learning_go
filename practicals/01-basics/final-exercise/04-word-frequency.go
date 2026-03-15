/*
In this question, we are building a tool that analyzes the frequency of words in a given text. You need to implement a function wordFrequency that receives a string and returns a map with the frequency of each word in the string.
package main

import (
    "fmt"
    "strings"
)

func wordFrequency(text string) map[string]int {
    // TODO: implement this function
}

func main() {
    text := "The quick brown fox jumps over the lazy dog"
    fmt.Println(wordFrequency(text))

}
*/

package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	// TODO: implement this function
	result := make(map[string]int)
	listOfString := strings.Split(text, " ")
	for _, value := range listOfString {
		for _, value_2 := range listOfString {
			if value == value_2 {
				result[value]++
			}
		}
	}
	return result
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))

}
