package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	doneChan := make(chan bool)
	defer close(doneChan)

	stream := generatorFunc[int, bool](doneChan, randomNumber)

	for val := range stream {
		fmt.Println(val)
	}
}

func generatorFunc[T any, V any](done <-chan V, fn func() T) <-chan T {
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

func randomNumber() int {
	maxValue := 1000000
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(maxValue)
}
