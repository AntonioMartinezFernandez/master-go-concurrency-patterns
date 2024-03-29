package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentProcessor(t *testing.T) {
	ctx := context.Background()

	processor := NewProcessor(10000, 2)
	go processor.Start(ctx)

	<-time.After(200 * time.Millisecond)

	assert.Equal(t, processor.elementsProcessed, 10000, "The number of elements processed must be 10000.")
	assert.Less(t, processor.elementsDeleted, 8000, "The number of elements deleted must be less than 8000.")
	assert.Greater(t, processor.elementsDeleted, 2000, "The number of elements deleted must be greater than 2000.")
}

func TestConcurrentProcessorWithStop(t *testing.T) {
	ctx := context.Background()

	processor := NewProcessor(10000, 2)
	go processor.Start(ctx)

	<-time.After(1 * time.Millisecond)
	processor.Stop()

	assert.Less(t, processor.elementsProcessed, 10000, "The number of elements processed must be less than 10000.")
	assert.Less(t, processor.elementsDeleted, 8000, "The number of elements deleted must be less than 8000.")
	assert.Greater(t, processor.elementsDeleted, 1, "The number of elements deleted must be greater than 1.")
}

func TestConcurrentProcessorWithCtxCancel(t *testing.T) {
	ctxWithCancel, cancelFunc := context.WithCancel(context.Background())

	processor := NewProcessor(10000, 2)
	go processor.Start(ctxWithCancel)

	<-time.After(1 * time.Millisecond)
	cancelFunc()

	assert.Less(t, processor.elementsProcessed, 10000, "The number of elements processed must be less than 10000.")
	assert.Less(t, processor.elementsDeleted, 8000, "The number of elements deleted must be less than 8000.")
	assert.Greater(t, processor.elementsDeleted, 1, "The number of elements deleted must be greater than 1.")
}
