package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func check(ctx context.Context, url string, wg *sync.WaitGroup, errCh chan error) {
	var errVal error

	defer func() {
		if errVal != nil {
			select {
			case errCh <- errVal:
			case <-ctx.Done():
				log.Println("aborting:", url)
			}
		}
		wg.Done()
	}()

	respond, ok := http.Get(url)
	if ok != nil {
		errVal = fmt.Errorf("error %w", ok)
		return
	}

	if respond.StatusCode != http.StatusOK {
		errVal = fmt.Errorf("wrong status: %s", url)
		return
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	errCh := make(chan error)
	urls := []string{
		"https://yahoo.com",
		"https://go0gle.com",
		"https://example.com",
	}
	for _, url := range urls {
		wg.Add(1)
		go check(ctx, url, wg, errCh)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	if err := <-errCh; err != nil {
		log.Println(err)
		cancel()
		return
	}

	log.Println("Succeful testing!")

}
