package main

import "sync"

func FanIn[T any](done <-chan bool, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	fannedInStream := make(chan T)

	for _, channel := range channels {
		wg.Add(1)
		go Transfer(&wg, done, fannedInStream, channel)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}

func Transfer[T any](wg *sync.WaitGroup, done <-chan bool, fannedInStream chan T, channel <-chan T) {
	defer wg.Done()

	for i := range channel {
		select {
		case <-done:
			return
		case fannedInStream <- i:
		}
	}
}
