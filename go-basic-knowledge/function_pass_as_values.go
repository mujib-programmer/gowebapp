// Functions are also variables in Go, we can use type to define them.
// Functions that have the same signature can be seen as the same type.
// What's the advantage of this feature? The answer is that it allows us to pass functions as values.

package main

import "fmt"

type testInt func(int) bool // define a function type of variable

func isOdd(integer int) bool {
	return integer%2 != 0
}

func isEven(integer int) bool {
	return integer%2 == 0
}

// pass the function f as an argument to another function
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

var slice = []int{1, 2, 3, 4, 5, 7}

func main() {
	odd := filter(slice, isOdd)
	even := filter(slice, isEven)

	fmt.Println("slice = ", slice)
	fmt.Println("Odd elements of slice are: ", odd)
	fmt.Println("Even elements of slice are: ", even)
}

/*
Expected output is:
slice =  [1 2 3 4 5 7]
Odd elements of slice are:  [1 3 5 7]
Even elements of slice are:  [2 4]
*/
