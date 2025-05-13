package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[key]++
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncre := func(key string, n int) {
		for i := 0; i < n; i++ {
			c.inc(key)
		}
		wg.Done()
	}

	wg.Add(3)

	go doIncre("a", 100)
	go doIncre("b", 100)
	go doIncre("a", 100)

	wg.Wait()
	fmt.Println(c.counters)

}
