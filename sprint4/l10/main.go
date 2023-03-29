package main

import "fmt"

func main() {
	channel := make(chan string, 1)
	channel <- "hello"
	close(channel)

	fmt.Println(<-channel)

	channel <- "hello"

	//fmt.Println(<-channel)

}
