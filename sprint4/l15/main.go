package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Counter struct {
	mu sync.Mutex
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	c.c[key]++
	c.mu.Unlock()
	wg.Done()
}
func (c *Counter) getValue(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c[key]
}

func main() {
	key := "test"
	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		c.Inc(key)
	}
	// time.Sleep(2 * time.Second)
	wg.Wait()
	fmt.Println(c.getValue(key))
}
