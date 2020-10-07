package main

import (
	"fmt"
)

func main() {
	a, b := 1, 3
	fmt.Println(add(&a, &b))

	printSlice("printSlice:", 1, 2, 3, 4, 5)

	fmt.Println(divide(0, 0))

	var d func(string) = func(msg string) { //anonymous function
		fmt.Println(msg)
	}
	d("anonumous function")
}

//a's type obtained from b
//it is better to send pointer for performance benefit
//it is also possible to return pointer
func add(a, b *int) int {
	return *a + *b
}

//values must be the last parameter
func printSlice(message string, values ...int) {
	fmt.Println(message, values)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
