package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) //struct{} is requires 0 memory. Signal only channel

func main() {
	go logger()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}

	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
	time.Sleep(100 * time.Millisecond)
}

func logger() {
	for {
		select { //work like switch but only with channels. stays blocked until a message is received on one channel
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %v\n", entry.time, entry.severity, entry.message)
		case <-doneCh:
			fmt.Println("\nShutting down logger")
			break
		}
	}
}
