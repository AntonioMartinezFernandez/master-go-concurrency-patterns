package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go someAsyncFunc(1)
	go someAsyncFunc(2)
	go someAsyncFunc(3)

	time.Sleep(1 * time.Second)

	fmt.Println("hi")
}

func someAsyncFunc(number int64) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("number:", strconv.FormatInt(number, 10))
}
