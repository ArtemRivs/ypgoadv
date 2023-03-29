package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	log.Println("Start")

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("job:", i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	log.Println("Finish")

}
