// One thing that Go is better at than C is that it supports multi-value returns.

package main

import "fmt"

// return results of A + B and A * B
func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func main() {
	x := 3
	y := 4

	xPlusY, xTimesY := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPlusY)
	fmt.Printf("%d * %d = %d\n", x, y, xTimesY)
}

/*
The output of this code will be:
3 + 4 = 7
3 * 4 = 12
*/
