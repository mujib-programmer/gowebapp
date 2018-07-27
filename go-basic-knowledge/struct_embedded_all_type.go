// struct_embedded_all_type.go
// All the types in Go can be used as embedded fields.

package main

import (
	"fmt"
)

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human     // struct as embedded field
	Skills    // string slice as embedded field
	int       // build-in type as embedded field
	specialty string
}

func main() {
	// initialize Student Jane
	jane := Student{Human: Human{"Jane", 35, 100}, specialty: "Biology"}

	// access fields
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her specialty is ", jane.specialty)

	// modify value of skill field
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)

	// modify embedded field
	jane.int = 3
	fmt.Println("Her preferred number is ", jane.int)

}

// In the above example, we can see that all types can be embedded fields
// and we can use functions to operate on them.

/*
Expected output is:
Her name is  Jane
Her age is  35
Her weight is  100
Her specialty is  Biology
Her skills are  [anatomy]
She acquired two new ones
Her skills now are  [anatomy physics golang]
Her preferred number is  3
*/
