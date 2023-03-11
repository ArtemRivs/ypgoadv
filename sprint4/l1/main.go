package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	const n = 5
	for i := 0; i < n; i++ {
		wg.Add(1)

		go func(i int) {
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("hi %d\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("The End")

}
