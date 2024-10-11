// Compare struct

package main

import "fmt"

type s1 struct {
	x int
}
type s2 struct {
	x int
}

func main() {
	c := s1{x: 5}
	c1 := s2{x: 5}

	c2 := s1{x: 6}
	c3 := s1{x: 5}

	if c != c2 {
		fmt.Println("c and c2 are different values")
	}
	if c == c3 {
		fmt.Println("c and c3 are same")
	}
	fmt.Println(c1)

}
