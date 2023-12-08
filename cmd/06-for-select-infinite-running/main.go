package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			select {
			default:
				fmt.Println("Doing things...")
			}
		}
	}()

	time.Sleep(1 * time.Second)
}
