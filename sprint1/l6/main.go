package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	for i := 0; i < 9999; i++ {

	}
	stop := time.Now()

	fmt.Println(stop.Sub(start))

}
