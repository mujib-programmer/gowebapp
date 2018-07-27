// oo_method_overriding.go
// If we want Employee to have its own method SayHi ,
// we can define a method that has the same name in Employee,
// and it will hide SayHi in Human when we call it.

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
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// this method override SayHi() method from Human
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) // yes you can split into 2 lines here
}

func main() {
	sam := Employee{Human{"Sam", 45, "111-888-xxxx"}, "Golang Inc"}
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}

	sam.SayHi()
	mark.SayHi()
}
