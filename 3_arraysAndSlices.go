package main

import (
	"fmt"
)

func main() {
	//arrays. [size]type{...}
	var g [2]int
	g[0] = 2
	grades := [...]int{1, 2, 3}
	fmt.Println("\nArray:", g, grades, len(grades))
	identityMatrix := [...][3]int{[...]int{1, 0, 0}, [...]int{0, 1, 0}, [...]int{0, 0, 1}}
	fmt.Println("identityMatrix:", identityMatrix)

	//arrys are copied not referenced
	a4 := [...]int{1, 2, 3}
	b4 := a4 //deep copy
	b4[1] = 5
	fmt.Println("Deep copied Arrays:", a4, b4)
	c4 := &a4 //reference
	c4[1] = 5
	fmt.Printf("Array Reference: %v, %v, %T, %T", a4, *c4, a4, c4)

	///////////////////////////////////////////////////////////////////////////////////////////////

	//slices
	a5 := []int{1, 2, 3}
	b5 := a5 //reference
	b5[1] = 5
	fmt.Println("\n\nSlices:", a5, b5, len(b5), cap(b5))

	c5 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(
		"Slices:",
		c5[:],   //slice of all elements
		c5[1:],  //slice from second element to end
		c5[:2],  //slice first 2 element
		c5[0:2], //slice 0, 1 element
	)
	//above operations can be done with slices

	d5 := make([]int, 3, 6)
	fmt.Println("Slices make:", d5, len(d5), cap(d5))
	d5 = append(d5, 1, 2)
	fmt.Println("Slices make:", d5, len(d5), cap(d5))
	d5 = append(d5, 3)
	fmt.Println("Slices make:", d5, len(d5), cap(d5))
	d5 = append(d5, 4)
	fmt.Println("Slices make:", d5, len(d5), cap(d5))
	d5 = append(d5, []int{5, 6, 7}...)
	fmt.Println("Slices make:", d5, len(d5), cap(d5))
	d5Slice := append(d5[0:4], d5[5:]...) //remove item from 4th index
	fmt.Println(
		"Remove middle element:",
		d5, //also changed
		d5Slice,
		len(d5Slice),
		cap(d5Slice),
		"\n",
	)
}
