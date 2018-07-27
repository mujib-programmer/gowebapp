// struct_overload_field.go
// There is one more problem however.
// If Human has a field called phone and Student has a field with same name,
// what should we do?
// Go use a very simple way to solve it. The outer fields get upper access levels,
// which means when you access student.phone , we will get the field called phone in student,
// not the one in the Human struct. This feature can be simply seen as field overload ing

package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string // Human as phone field
}

type Employee struct {
	Human
	specialty string
	phone     string // phone in employee
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-xxxx"}, "Designer", "333-222"}

	fmt.Println("Bob's work phone is:", Bob.phone)
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)

}

/*
Expected output is:
Bob's work phone is: 333-222
Bob's personal phone is: 777-444-xxxx
*/
