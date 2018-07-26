package main

import "fmt"

// simple function to add 1 to a
func add1(a *int) int {
	*a = *a + 1 // we changed value of a
	return *a
}

func main() {
	x := 3

	fmt.Println("x = ", x) // should print "x = 3"

	x1 := add1(&x) // call add1(&x) pass memory address of x

	fmt.Println("x+1 = ", x1) // should print "x+1 = 4"
	fmt.Println("x = ", x)    // should print "x = 4"

}

/*
the output of this code should be:
x =  3
x+1 =  4
x =  4
*/

// Now we can change the value of x in the functions. Why do we use pointers? What are the advantages?
// * Allows us to use more functions to operate on one variable.
// * Low cost by passing memory addresses (8 bytes), copy is not an efficient way, both in terms of time and space, to pass variables.
// * channel , slice and map are reference types, so they use pointers when passing to functions by default. (Attention:
//   If you need to change the length of slice , you have to pass pointers explicitly)
