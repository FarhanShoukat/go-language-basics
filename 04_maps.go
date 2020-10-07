package main

import (
	"fmt"
)

func main() {
	//maps. has no order
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	m["c"] = 3
	delete(m, "b")
	_, maOK := m["a"]
	md, mdOK := m["d"]
	fmt.Println("Map:", m, m["a"], maOK, md, mdOK, len(m))
}
