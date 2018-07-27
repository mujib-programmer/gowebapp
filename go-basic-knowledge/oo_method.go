// oo_method.go
// In the example above, the Area() methods belong to both Rectangle and Circle respectively,
// so the receivers are Rectangle and Circle.
// One thing that's worth noting is that the method with a dotted line means
// the receiver is passed by value, not by reference.
// The difference between them is that a method can change its receiver's values
// when the receiver is passed by reference, and it gets a copy of the receiver
// when the receiver is passed by value.

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

// method
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

// method
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	c1 := Circle{10}
	c2 := Circle{25}
	r1 := Rectangle{9, 4}
	r2 := Rectangle{12, 2}

	fmt.Println("Area of c1 is: ", c1.Area())
	fmt.Println("Area of c2 is: ", c2.Area())
	fmt.Println("Area of r1 is: ", r1.Area())
	fmt.Println("Area of r2 is: ", r2.Area())
}

// Notes for using methods:
// * If the name of methods are the same but they don't share the same receivers,
//   they are not the same.
// * Methods are able to access fields within receivers.
// * Use . to call a method in the struct, the same way fields are called.

/*
Expected result is:
Area of c1 is:  314.1592653589793
Area of c2 is:  1963.4954084936207
Area of r1 is:  36
Area of r2 is:  24
*/
