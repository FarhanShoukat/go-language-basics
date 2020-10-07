package main

import "fmt"

func main() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	i := 0
	for ; i < 5; i = i + 1 { //both initializer and increment are operational
		fmt.Println(i)
	}

	i = 0
	for i < 5 { //while loop
		fmt.Println(i)
		i++
	}

	fmt.Println("\nInfinite loop")
	i = 0
	for { //infinite loop
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	//continue is here as well. nested loops are here as well like c++
	fmt.Println("\nLoop")
Loop:
	for i = 1; i < 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	s := []int{1, 2, 3}
	fmt.Println("Collections:", s)
	for k, v := range s {
		fmt.Println(k, v)
	}
}
