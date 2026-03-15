// These are functions that receives function as an arg or return function as a output
/*
// why to use it ?
-  Composition
	- creating smaller function that take care of certain piece of logic
	- Composing complex function by using different smaller function
- Reduced bugs
- Code gets easier to read and understand

Where it is used? Eg:
*/

package main

import "fmt"

// Normal approach
func calcArea(r float64) float64 {
	return 3.14 * r * r
}

func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}

func calcDiameter(r float64) float64 {
	return 2 * r
}

// High Order Approach
func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Thankyou!")
}
func getFunction(query int) func(r float64) float64 {
	query_to_func := map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_func[query]
}

func main() {

	//Normal Apprach
	var query int
	var radius float64

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)
	fmt.Printf("Enter \n 1 - Area, \n 2 - perimeter \n 3 - diameter:")
	fmt.Scanf("%d", &query)
	if query == 1 {
		fmt.Println("Result: ", calcArea(radius))
	} else if query == 2 {
		fmt.Println("Result Perimeter: ", calcPerimeter(radius))
	} else if query == 3 {
		fmt.Println("Result Dia:", calcDiameter(radius))
	} else {
		fmt.Println("Invalid Query")
	}

	// High Order Approach
	printResult(radius, getFunction(query))
}
