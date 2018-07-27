// oo_method_inheritance.go
// We learned about inheritance of fields in the last section.
// Similarly, we also have method inheritance in Go.
// If an anonymous field has methods, then the struct that contains the field
// will have all the methods from it as well.

package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // anonymous field
	school string
}

type Employee struct {
	Human
	company string
}

// define a method in Human
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	sam := Employee{Human{"Sam", 45, "111-888-xxxx"}, "Golang inc"}
	mark := Student{Human{"Mark", 25, "222-222-yyyy"}, "MIT"}

	sam.SayHi()
	mark.SayHi()
}

/*
Expected output is:
Hi, I am Sam you can call me on 111-888-xxxx
Hi, I am Mark you can call me on 222-222-yyyy
*/
