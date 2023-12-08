package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Create channels (buffered)
	inputChannel := make(chan string, 1)
	outputChannel := make(chan string, 1)

	// Init 'displayer' and 'processor' go routines
	go dataDisplayer(outputChannel)
	go dataProcessor(inputChannel, outputChannel)

	// Start 'data generator' go routine
	for i := 0; i < 10; i++ {
		go dataGenerator(inputChannel, strconv.Itoa(i)+" Hello, channels!")
	}

	// Wait 1 second...
	time.Sleep(1 * time.Second)
}

func dataGenerator(inputChannel chan string, data string) {
	// Just publish the data to the input channel
	inputChannel <- data
}

func dataProcessor(inputChannel chan string, outputChannel chan string) {
	// Receive data from input channel, process data, publish processed data to the output channel (indefinitely)
	for {
		receivedData := <-inputChannel
		outputChannel <- receivedData + " - processed OK"
	}
}

func dataDisplayer(outputChannel chan string) {
	// Receive and print data from output channel (indefinitely)
	for {
		processedData := <-outputChannel
		fmt.Println(processedData)
	}
}
