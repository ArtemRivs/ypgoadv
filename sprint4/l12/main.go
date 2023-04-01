package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		fmt.Println("wait")
		select {
		case <-ctx.Done():
			fmt.Println("cancel")
			return
		default:
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(time.Second)

	cancel()

	time.Sleep(time.Second)

}
