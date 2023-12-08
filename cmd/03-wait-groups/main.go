package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var myWaitGroup sync.WaitGroup

	for i := 1; i <= 5; i++ {
		myWaitGroup.Add(1)

		i := i //! https://go.dev/doc/faq#closures_and_goroutines

		go func() {
			defer myWaitGroup.Done()
			worker(i)
		}()
	}

	myWaitGroup.Wait()

	fmt.Println("Finish!")
}

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(1 * time.Second)

	fmt.Printf("Worker %d done\n", id)
}
