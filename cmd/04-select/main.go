package main

import (
	"fmt"
)

func main() {
	channelOne := make(chan string, 1)
	channelTwo := make(chan string, 1)

	go func() {
		channelOne <- "cat"
	}()

	go func() {
		channelTwo <- "dog"
	}()

	//! Select statement will block the execution until it receive data
	select {
	case msgFromChannelOne := <-channelOne:
		fmt.Println(msgFromChannelOne)
	case msgFromChannelTwo := <-channelTwo:
		fmt.Println(msgFromChannelTwo)
	}
}
