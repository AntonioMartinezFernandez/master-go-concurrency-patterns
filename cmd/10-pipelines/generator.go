package main

import (
	"math/rand"
	"time"
)

func CreateStreamFromGenerator[T any](done <-chan bool, fn func() T) <-chan T {
	//! IMPORTANT: 'stream' is an unbuffered channel because it will block the sending go routine (the function under this comment)
	//! until receiving go routine (the TakeValues go routine function) has finished processing each value.
	stream := make(chan T)

	go func() {
		defer close(stream)

		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func GenerateRandomInt() int {
	maxValue := 500_000_000
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(maxValue)
}
