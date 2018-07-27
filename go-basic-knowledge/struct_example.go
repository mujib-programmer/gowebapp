// struct_example.go
package main

import (
	"fmt"
)

// define a new type
type person struct {
	name string
	age  int
}

// struct is passed by value
// compare the age of two people, then return the older person and difference of age
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person

	tom.name, tom.age = "tom", 18
	bob := person{age: 25, name: "Bob"}
	paul := person{"Paul", 43}

	tb_older, tb_diff := Older(tom, bob)
	tp_older, tp_diff := Older(tom, paul)
	bp_older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tb_older.name, tb_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, paul.name, tp_older.name, tp_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, paul.name, bp_older.name, bp_diff)
}

/*
Expected result is:
Of tom and Bob, Bob is older by 7 years
Of tom and Paul, Paul is older by 25 years
Of Bob and Paul, Paul is older by 18 years
*/
