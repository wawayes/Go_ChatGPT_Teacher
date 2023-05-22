package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	counter int
	mu      sync.Mutex
}

func (sc *SafeCounter) Inc() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.counter++
}

func (sc *SafeCounter) Value() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.counter
}

func demo005() {
	wg := sync.WaitGroup{}
	counter := SafeCounter{counter: 0}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Final Count: %d", counter.Value())
}
