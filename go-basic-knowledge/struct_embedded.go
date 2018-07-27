// struct_embedded.go
package main

import (
	"fmt"
)

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human     // embedded field, it means Student struct include all fields that human has
	specialty string
}

func main() {
	// instantiate and initialize a student
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	// access fields
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His specialty is ", mark.specialty)

	// modify mark specialty
	mark.specialty = "AI"
	fmt.Println("Mark changed his specialty")
	fmt.Println("His specialty is ", mark.specialty)

	fmt.Println("Mark become old. He is not an athlete anymore")
	mark.age = 46
	mark.weight += 60
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is", mark.weight)

}

/*
Expected output is:
His name is  Mark
His age is  25
His weight is  120
His specialty is  Computer Science
Mark changed his specialty
His specialty is  AI
Mark become old. He is not an athlete anymore
His age is  46
His weight is 180
*/
