package main

func TakeValues[T any](done <-chan bool, stream <-chan T, items int) <-chan T {
	takenValues := make(chan T)

	go func() {
		defer close(takenValues)

		for i := 0; i < items; i++ {
			select {
			case <-done:
				return
			case takenValues <- <-stream:
				//! Why 2 arrows? We READ FROM 'stream' channel, and at the same time we WRITE TO 'takenValues' channel.
				//! Is the equivalent of assign a value from 'stream' to a variable, and write this value to 'takenValues'
			}
		}
	}()

	return takenValues
}
