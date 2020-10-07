package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) //channel that can store 50 integers. optional
	wg.Add(2)

	go func(ch <-chan int /*receive only channel*/) {
		//i := <-ch	//extract message from channel
		//i, ok := <-ch	//extract message from channel with flag ok telling if channel is open
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int /*send only channel*/) {
		ch <- 42 //push message to channel
		ch <- 50
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
