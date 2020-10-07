package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "io/ioutil"
	// "log"
	// "net/http"
)

func main() {
	a, b := "start", "b"
	defer fmt.Println(a) //runs after the end of function and in reverse order
	defer fmt.Println(b)
	a = "end"

	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() //runs after the end of function

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
