package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	//type conversion
	consoleWriter, ok := w.(ConsoleWriter)
	fmt.Println(consoleWriter, ok)
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

//implicit implementation of Writer.Write for ConsoleWriter
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

/////////////////////////////////////////////////////////////////
//interface composition
type Closer interface {
	Close()
}

type WriterCloser interface { //similar to embedding in structs
	Writer
	Closer
}
