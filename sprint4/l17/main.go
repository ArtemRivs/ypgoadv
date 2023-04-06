package main // объявление пакета main

import (
	"fmt"
)

func main() {
	valueCh := make(chan int)
	quitCh := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-valueCh)
		}
		quitCh <- true
	}()

	processValue(valueCh, quitCh)

}

func processValue(valueCh chan int, quitCh chan bool) {
	x := 1
	for {
		select {
		case valueCh <- x:
			x = x * 2
		case <-quitCh:
			return
		}
	}
}
