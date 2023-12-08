package main

import (
	"fmt"
	"time"
)

func main() {
	doneChan := make(chan bool)

	go infiniteFunction(doneChan)

	time.Sleep(1 * time.Second)
	close(doneChan) // Anoother option: doneChan <- true

	time.Sleep(2 * time.Second)
}

func infiniteFunction(doneChan <-chan bool) { //! The '<-' syntax, stablish the 'doneChan' channel as read-only
	for {
		select {
		case <-doneChan:
			return
		default:
			fmt.Println("Doing things...")
		}
	}
}
