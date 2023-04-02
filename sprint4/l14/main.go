package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	c.c[key]++
	c.mu.Unlock()
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
		go c.Inc(key)
	}
	// time.Sleep(2 * time.Second)

	fmt.Println(c.getValue(key))
}
