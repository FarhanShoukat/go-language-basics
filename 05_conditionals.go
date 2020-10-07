package main

import "fmt"

func main() {
	//operators: <, >, ==, !=, ||, &&, !
	m := map[string]int{"a": 1, "b": 2}
	if pop, ok := m["a"]; ok {
		fmt.Println(pop)
	}

	x := 0
	if x == 0 {
		fmt.Println(0)
	} else if x == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}

	switch i := 1 + 1; i {
	case 1, 3, 5:
		fmt.Println(1, 3, 5)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println(-1)
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("i <= 10")
		fallthrough
	case i >= 20:
		fmt.Println("i >= 20")
	default:
		fmt.Println("> 20")
	}

	fmt.Println() /////////////////////////////////////////////////
	//type switches
	var i1 interface{} = 1
	switch i1.(type) {
	case int:
		fmt.Println("i1 is an int")
		break
		fmt.Println("this will print too")
	case float64:
		fmt.Println("i1 is a float64")
	}
}
