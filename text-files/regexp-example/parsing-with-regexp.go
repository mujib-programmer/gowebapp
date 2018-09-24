package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	// Find the first match
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	// Find all matches and save to a slice, n less then 0 means return all matches, indicates length of slice if it's greater than 0.
	all := re.FindAll([]byte(a), -1)
	fmt.Println("Find", all)

	// Find index of first match, start and end position.
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	// Find index of all matches, the n does same job above.
	allIndex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex:", allIndex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	// Find first submatch and return array, the first element contains all elements, the second element contains the result of first (), the third element contains the result of second ().
	// Output:
	// the first element: "am learning Go language"
	// the second element: " learning Go ", notice spaces will be outputed as well.
	// the third element: "uage"
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch:", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	// Same as FindIndex().
	submatchIndex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchIndex)

	// FindAllSubmatch, find all submatches.
	submatchAll := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchAll)

	// FindAllSubmatchIndex, find index of all submatches.
	submatchAllIndex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchAllIndex)

}
