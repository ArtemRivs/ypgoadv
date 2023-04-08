package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func healthCheck(ctx context.Context, url string, errCh chan<- error, wg *sync.WaitGroup) {
	var defErr error
	defer func() {
		if defErr != nil {
			select {
			case errCh <- defErr:
			case <-ctx.Done():
				log.Println("aborting", url)
			}
		}
		wg.Done()
	}()

	resp, err := http.Get(url)
	if err != nil {
		defErr = fmt.Errorf("healthcheck failed: %w", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		defErr = errors.New("healthcheck failed: status not ok")
		return
	}
}

func main() {
	wg := &sync.WaitGroup{}
	errCh := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	hostsToCheck := []string{
		"https://yandex.ru",
		"https://google.com",
		"https://bing.com",
		"https://test000000001.com",
		"https://test000000002.com",
		"https://test000000003.com",
		"https://test000000004.com",
		"https://test000000005.com",
		"https://test000000006.com",
		"https://test000000007.com",
		"https://test000000008.com",
		"https://test000000009.com",
		"https://ya.ru",
		"https://lamoda.ru",
	}
	for _, hostToCheck := range hostsToCheck {
		log.Println("checking", hostToCheck)
		wg.Add(1)
		go healthCheck(ctx, hostToCheck, errCh, wg)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	if err := <-errCh; err != nil {
		log.Println(err)
		//time.Sleep(2 * time.Second)
		cancel()
		return
	}

	log.Println("successful healthcheck")
}
