package main

import "fmt"

func main() {
	// start - input
	data := []int{1, 2, 3, 4, 5}

	// stage 1
	numsChannel := sliceToChannel(data)

	// stage 2
	squaredChannel := square(numsChannel)

	// stage 3 - output
	for elem := range squaredChannel {
		fmt.Println(elem)
	}
}

func sliceToChannel(input []int) <-chan int {
	output := make(chan int) //! Unbuffered channel - Communication between stages will be synchronous

	go func() {
		for _, elem := range input {
			output <- elem
		}
		close(output)
	}()
	return output
}

func square(input <-chan int) <-chan int {
	output := make(chan int) //! Unbuffered channel - Communication between stages will be synchronous

	go func() {
		for elem := range input {
			output <- elem * elem
		}
		close(output)
	}()

	return output
}
