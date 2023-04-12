package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func healthCheck(url string, errCh chan<- error, wg *sync.WaitGroup, stopCh <-chan struct{}) {
	var defErr error
	defer func() {
		if defErr != nil {
			select {
			case errCh <- defErr: // первая горутина, поймавшая ошибку, сможет записать в канал
				log.Println("send error to error channel")
			case <-stopCh: // остальные завершат работу, провалившись в этот case
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
	stopCh := make(chan struct{})
	hostsToCheck := []string{
		"https://g00gle.com",
		"https://test000000001.com",
		"https://yahoo222.com",
		"https://example.com",
	}
	for _, hostToCheck := range hostsToCheck {
		log.Println("checking", hostToCheck)
		wg.Add(1)
		go healthCheck(hostToCheck, errCh, wg, stopCh)
	}
	// в отдельной горутине ждём завершения всех healthCheck после этого закрываем канал errCh — больше записей не будет
	go func() {
		wg.Wait()
		log.Println("wait group!")
		close(errCh)
	}()

	if err := <-errCh; err != nil {
		log.Println(err)
		close(stopCh)
		return
	}

	log.Println("successful healthcheck")
}
