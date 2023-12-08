package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	timeStart := time.Now()

	doneChan := make(chan bool)
	defer close(doneChan)

	//! Naive/Slow solution
	// rndIntStream := CreateStreamFromGenerator(doneChan, GenerateRandomInt) // rndIntStream := generatorFunc[int](doneChan, generateRandomInt)
	// primeNumbers := PrimeFinder(doneChan, rndIntStream)
	// takenValues := TakeValues(doneChan, primeNumbers, 10) // takenValues := takeValues[int](doneChan, primeNumbers, 10)
	//
	// for val := range takenValues {
	// 	fmt.Println(val)
	// }

	//! Concurrent solution

	// Fan Out
	CPUcount := runtime.NumCPU() // CPUs of the machine where the application is running
	fmt.Println("Number of CPUs: ", CPUcount)

	primeFinderChannels := make([]<-chan int, CPUcount) // Slice of read-only channels, with a length equal to the number of CPUs

	rndIntStream := CreateStreamFromGenerator(doneChan, GenerateRandomInt) // Create the stream of random numbers

	for i := 0; i < CPUcount; i++ {
		primeFinderChannels[i] = PrimeFinder(doneChan, rndIntStream) // Create one prime finder instance (and channel) per CPU thread
	}

	// Fan In
	fannedInStream := FanIn(doneChan, primeFinderChannels...)

	takenValues := TakeValues(doneChan, fannedInStream, 10)

	for value := range takenValues {
		fmt.Println(value)
	}

	fmt.Println("Total time: ", time.Since(timeStart).Milliseconds(), "ms")
}
