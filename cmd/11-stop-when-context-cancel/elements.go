package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type Element struct {
	id     string
	name   string
	delete bool
}

type Processor struct {
	mut                  sync.Mutex
	wg                   sync.WaitGroup
	numberOfElements     int
	rawElementsChan      chan Element
	elementsToDeleteChan chan Element
	cancelFunc           context.CancelFunc
	elementsProcessed    int
	elementsDeleted      int
}

func NewProcessor(elems int) *Processor {
	return &Processor{
		mut:                  sync.Mutex{},
		wg:                   sync.WaitGroup{},
		numberOfElements:     elems,
		rawElementsChan:      make(chan Element),
		elementsToDeleteChan: make(chan Element),
		elementsProcessed:    0,
		elementsDeleted:      0,
	}
}

func (p *Processor) Start(ctx context.Context) {
	ctx, cancelFunc := context.WithCancel(ctx)
	p.cancelFunc = cancelFunc
	defer p.Stop()

	fmt.Println(">>> Starting process")

	p.wg.Add(1)
	go p.elementsDeleter(ctx)
	p.wg.Add(1)
	go p.elementsFilter()
	p.wg.Add(1)
	go p.elementsGenerator()

	p.wg.Wait()
}

func (p *Processor) Stop() {
	fmt.Println(">>> Stopping process")
	p.cancelFunc()
}

func (p *Processor) elementsGenerator() {
	defer func() {
		p.wg.Done()
		close(p.rawElementsChan)
	}()

	for i := 0; i < p.numberOfElements; i++ {
		elem := Element{
			id:     strconv.Itoa(i),
			name:   strconv.Itoa(i),
			delete: rand.Intn(2) == 1,
		}

		p.rawElementsChan <- elem
	}
}

func (p *Processor) elementsFilter() {
	defer func() {
		p.wg.Done()
		close(p.elementsToDeleteChan)
	}()

	for elem := range p.rawElementsChan {
		p.mut.Lock()
		p.elementsProcessed++
		p.mut.Unlock()

		if elem.delete {
			p.elementsToDeleteChan <- elem
		}
	}
}

func (p *Processor) elementsDeleter(ctx context.Context) {
	defer p.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case elem, ok := <-p.elementsToDeleteChan:
			if !ok {
				return
			}
			p.mut.Lock()
			p.elementsDeleted++
			p.mut.Unlock()

			fmt.Printf("deleting element %s\n", elem.id)
		}
	}
}
