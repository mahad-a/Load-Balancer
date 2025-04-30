package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	targetURL     = "http://localhost:8080"
	requestRate   = 10 * time.Millisecond
	numGoroutines = 10
	testDuration  = 10 * time.Second
)

func sendRequests(wg *sync.WaitGroup, stopCh <-chan struct{}) {
	defer wg.Done()
	client := http.Client{}
	for {
		select {
		case <-stopCh:
			return
		default:
			start := time.Now()
			resp, err := client.Get(targetURL)
			duration := time.Since(start)

			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				resp.Body.Close()
				fmt.Printf("Response in %v ms\n", duration.Milliseconds())
			}
			time.Sleep(requestRate)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	stopCh := make(chan struct{})

	// Launch multiple goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go sendRequests(&wg, stopCh)
	}

	// Let the test run for a fixed duration
	time.Sleep(testDuration)
	close(stopCh)

	wg.Wait()
	fmt.Println("Load test complete.")
}
