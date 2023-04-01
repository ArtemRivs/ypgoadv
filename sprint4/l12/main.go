package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancel")
			return
		default:
			i++
			fmt.Println("wait:", i)
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(1000 * time.Millisecond)

	cancel()

	time.Sleep(time.Second)

}
