package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()

	processor := NewProcessor(100000, 50)
	go processor.Start(ctx)

	<-time.After(5 * time.Millisecond)
	processor.Stop()

	<-time.After(2 * time.Second)
}
