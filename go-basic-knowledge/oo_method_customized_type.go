// oo_method_customized_type.go
// Can the receiver only be a struct? Of course not. Any type can be the receiver of a method.
// You may be confused about customized types.
// Struct is a special kind of customized type -there are more customized types.
// Use the following format to define a customized type.
//    `type typeName typeLiteral`
// I hope that you know how to use customized types now. Similar to typedef in C.

package main

import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

// Define a struct Box which has fields height, width, length and color.
type Box struct {
	width, height, depth float64
	color                Color
}

// Use Color as alias of byte
type Color byte

// Define a struct BoxList which has Box as its field
type BoxList []Box // a slice of boxes

// Volume() uses Box as its receiver and returns the volume of Box
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// method with a pointer receiver
// SetColor( c Color) changes Box's color.
// Its receiver is a pointer of Box. Yes, you can use *Box as a receiver.
// Why do we use a pointer here? Because we want to change Box's color in this method.
// Thus, if we don't use a pointer, it will only change the value inside a copy of Box.
func (b *Box) setColor(c Color) {

	// You might be asking why we aren't using (*b).Color=c instead of b.Color=c in the SetColor() method.
	// Either one is OK here because Go knows how to interpret the assignment.
	// Do you think Go is more fascinating now?
	b.color = c
}

// BiggestsColor() returns the color which has the biggest volume.
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

// PaintItBlack() sets color for all Box in BoxList to black.
func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {

		// You may also be asking whether we should use (&bl[i]).SetColor(BLACK) in PaintItBlack
		// because we pass a pointer to SetColor .
		// Again, either one is OK because Go knows how to interpret it!
		bl[i].setColor(BLACK)
	}
}

// String() use Color as its receiver, returns the string format of color name.
func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm3")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	// Let's paint them all black
	boxes.PaintItBlack()

	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())

}

/*
Expected output is:
We have 6 boxes in our set
The volume of the first one is 64 cm3
The color of the last one is YELLOW
The biggest one is YELLOW
The color of the second one is BLACK
Obviously, now, the biggest one is BLACK
*/
