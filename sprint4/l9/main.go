package main

import "time"

func main() {
	bufSize := 10
	channel := make(chan int, bufSize)
	for i := 0; i < 20; i++ {
		go func(i int) {
			// time.Sleep(time.Millisecond)
			channel <- i
		}(i)
	}
	channel <- 1
	time.Sleep(time.Second)

}
