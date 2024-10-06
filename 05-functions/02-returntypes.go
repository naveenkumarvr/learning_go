package main

import "fmt"

func operations(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	// Two return value, which is also declared at beginning
	return sum, diff
}

// Alternative way with keyword for return, This is called Named Return Parameter
func operations_named(c int, d int) (sum int, diff int) {
	sum = c + d
	diff = c - d
	// Since the return variable are declared at the top not required to put it here again
	// one should not use any variable here as the name are returned at the beginning
	return
}

// Variadic Functions: Accepts variable number of arguments.
// Syntax: func <func_name> (parm1 type, parm-2 type, parm3 ...type) <return type>
// The above statement says that the function accepts 2 parameter of specific type and n number of parameter on last type
// Variadic parameter should be placed at the end
func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
func printDetails(student string, subjects ...string) {
	fmt.Println("Hey ", student, ", here are your subjects -")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
}

// Blank Identifier '_'
func f() (int, int) {
	return 42, 55
}

func main() {
	resultsum, resultdiff := operations(5, 2)
	fmt.Println(resultsum, resultdiff)

	//Named args
	resultsumn, resultdiffn := operations_named(5, 2)
	fmt.Println(resultsumn, resultdiffn)

	//Variadic Function
	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(10))
	fmt.Println(sumNumbers(10, 20))
	fmt.Println(sumNumbers(10, 20, 30))

	printDetails("Joe", "Physics", "Golang", "Python")

	//Blank Identifier
	a, b := f()
	fmt.Println("\n", a, b)

	// If we need only one var but function returns multiple we can use "_" to ignore the other variable
	c, _ := f()
	fmt.Println("\n", c)

}
