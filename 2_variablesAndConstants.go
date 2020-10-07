package main

import (
	"fmt"
)

func main() {
	//int, int8, int16, int32, int64, uint, float32, float64, bool, string, complex64, complex128
	//int is of varying depending on os
	//uint is unsigned integer
	//operations available for int are +, -, *, /, %
	//operations available for float are +, -, *, /. declaration 3.14, 13e18, 13E18, 13.7e12
	//cannot mix types in operation. e.g. int + uint is invalid
	a := 1 + 2i
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", real(a), real(a))
	fmt.Printf("%v, %T\n", imag(a), imag(a))
	fmt.Println(complex(1, 2))

	//strings are collection utf-8 characters. immutable
	s := "string1"
	fmt.Printf("\n%v, %v, %T\n", s[2], string(s[2]), s[2])
	s2 := "string2"
	fmt.Printf("%v, %T\n", s+s2, s+s2)
	b := []byte(s)
	fmt.Printf("%v, %T\n\n", b, b)

	//rune is utf-32 chracter. are basically int32
	r := 'a'
	fmt.Printf("%v, %T\n\n", r, r)

	//constants must be defined at compile time. e.g. a function call is invalid
	//collections like arrays are always variables
	const myConst = 1
	const c = 3.4
	fmt.Printf("%v, %T\n\n", myConst+c, myConst+c)

	//enumerations
	const (
		_ = iota + 10
		a1
		b1
		c1
	)
	fmt.Println("Enumerations:", a1, b1, c1)
	const (
		a2 = iota
	)
	fmt.Println("Enumerations Reset:", a2)
	//in following every variable occupies 1-bit. allows to encode roles in stream of bits
	const (
		a3 = 1 << iota
		b3
		c3
		d3
	)
	fmt.Println("Enumnerations Bit shifted:", a3, b3, c3, d3)
	var roles = a3 | b3 | c3
	fmt.Println("Roles:", roles)
	fmt.Println("Is a3: ", a3&roles == a3)
	fmt.Println("Is d3: ", d3&roles == d3)
}
