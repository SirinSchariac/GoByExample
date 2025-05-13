package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func worker(ops *atomic.Uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		ops.Add(1)
	}
}

func main() {
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(&ops, &wg)
	}
	wg.Wait()
	fmt.Println("ops:", ops.Load())
}
