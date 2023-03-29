package main

import (
	"fmt"
	"time"
)

func sender(c chan<- string) {
	c <- "Hello"
}

func resiver(c <-chan string) {
	// msg := <-c
	fmt.Printf("%s world!\n", <-c)
}

func main() {

	channel := make(chan string)
	go sender(channel)
	go resiver(channel)

	time.Sleep(time.Second)

}
