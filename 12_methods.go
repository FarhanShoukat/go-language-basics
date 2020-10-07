package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "hello",
		name:     "Go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() { //g passed by reference. remove * to pass by value
	fmt.Println(g.greeting, g.name)
}
