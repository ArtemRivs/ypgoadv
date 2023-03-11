package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	n := 10
	wg.Add(10)
	for i := 0; i < n-1; i++ {
		go func(v int) {
			fmt.Println(v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
