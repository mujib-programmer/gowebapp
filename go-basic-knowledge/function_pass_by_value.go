// When we pass an argument to the function that was called,
// that function actually gets the copy of our variables so any change
// will not affect to the original variable.

package main

import "fmt"

// simple function to add 1 to a
func add1(a int) int {
	a = a + 1 // we cange value of a
	return a  // return new value of a
}

func main() {
	x := 3

	fmt.Println("x = ", x) // should print "x = 3"

	x1 := add1(x) // call add1(x)

	fmt.Println("x+1 = ", x1) // should print "x+1 = 4"
	fmt.Println("x = ", x)    // should print "x = 3"

}

/*
the output should be:
x =  3
x+1 =  4
x =  3
*/

// Can you see that? Even though we called add1 with x , the origin value of x doesn't change.
// The reason is very simple: when we called add1 , we gave a copy of x to it, not the x itself.
// Now you may ask how I can pass the real x to the function.
// We need use pointers here.
