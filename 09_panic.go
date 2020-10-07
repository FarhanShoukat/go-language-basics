package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	//if panic happens, then defer is called first then panic

	fmt.Println("about to panic")

	defer func() {
		if err := recover(); err != nil { //catch.recover is always called in defer
			log.Println("Error:", err)
			//panic(err)
		}
	}()
	panic("something bad happened") //throw exception
	fmt.Println("panicker end")
}
