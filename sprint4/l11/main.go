package main

import "fmt"

func main() {
	firstChannel := make(chan int)
	secondChannel := make(chan int)

	select {
	case msg := <-firstChannel:
		fmt.Println(msg)
	case msg := <-secondChannel:
		fmt.Println(msg)
	default:
		fmt.Println("default action")
	}

}
