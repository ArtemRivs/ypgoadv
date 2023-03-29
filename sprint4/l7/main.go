package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	log.Println("Start main")

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(1 * time.Second)
			fmt.Println("job:", i)
		}(i)
	}

	time.Sleep(1 * time.Second)
	log.Println("Finish main")
	time.Sleep(1 * time.Second)

}
