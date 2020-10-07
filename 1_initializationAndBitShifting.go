package main

import (
	"fmt"
	"math"
)

func main() {
	var x int //by default variables have 0 value. bools will be false
	x = 0
	var y int = 0
	z := 0 //syntax only allowed in functions
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n\n", z, z)

	fmt.Println(math.Abs(math.Pow(math.Sqrt(10), 2)))

	//bit shifting
	fmt.Println("Bit shifting", 64<<2, 64>>2)
}
