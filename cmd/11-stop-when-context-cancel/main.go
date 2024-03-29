package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()

	processor := NewProcessor(10000)
	go processor.Start(ctx)

	<-time.After(1 * time.Millisecond)
	processor.Stop()

	<-time.After(2 * time.Second)
}
