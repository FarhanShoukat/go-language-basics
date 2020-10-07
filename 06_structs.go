package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string `required max:"100"` //tags
	companions []string
}

type Animal struct {
	name string
}

type Bird struct {
	Animal
	CanFly bool
}

func main() {
	//structs are deeply copied like arrays. pointers can be used

	d := Doctor{
		number:    3,
		actorName: "actor",
		companions: []string{
			"first",
			"second",
		},
	}
	fmt.Println(d, d.actorName)

	d2 := Doctor{3, "actor", []string{"first"}}
	fmt.Println(d2)

	d3 := struct{ name string }{name: "actor"} //anonymous struct
	fmt.Println(d3)

	//embedding
	b := Bird{
		Animal: Animal{name: "test"},
		CanFly: true,
	}
	fmt.Println(b, b.name)

	//tags
	t := reflect.TypeOf(Doctor{})
	field, f := t.FieldByName("actorName")
	fmt.Println("Tag:", field, field.Tag, f)
}
