package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

// Human implements fmt.Stringer
func (h Human) String() string {
	return "Name:" + h.name + ", Age:" + strconv.Itoa(h.age) + " years, Contact:" + h.phone
}

func main() {
	Bob := Human{"Bob", 39, "000-7777-xxx"}

	// any type that implements interface Stringer can be passed to fmt.Println as an argument.
	fmt.Println("This Human is :", Bob)
}
