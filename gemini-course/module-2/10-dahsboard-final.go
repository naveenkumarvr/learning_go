package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Latency    time.Duration
	Err        error
}

func checkWebsite(url string, ch chan<- Result) {
	start := time.Now()

	// 1. Dail the target url using http reuqest
	resp, err := http.Get(url)

	// Gate 1: Check for connection error
	if err != nil {
		ch <- Result{
			URL:        url,
			StatusCode: 0,
			Latency:    time.Since(start),
			Err:        err,
		}
		return
	}

	// Gate 2 Safely Close Resp Body
	defer resp.Body.Close()

	// Gate3 Http Error  and SuccessPath
	ch <- Result{
		URL:        url,
		StatusCode: resp.StatusCode,
		Latency:    time.Since(start),
		Err:        nil,
	}
}

func main() {
	targets := []string{
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/500",
		"https://this-domain-does-not-exist-at-all-12345.com",
	}
	resultChan := make(chan Result)
	var wg sync.WaitGroup

	for _, url := range targets {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			checkWebsite(u, resultChan)
		}(url)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	fmt.Println("==================================================")
	fmt.Printf("%-45s | %-6s | %-10s\n", "TARGET URL", "STATUS", "LATENCY")
	fmt.Println("==================================================")

	for res := range resultChan {
		if res.Err != nil {
			// Print error status formatted neatly inside the table columns
			fmt.Printf("%-45s | %-6s | %-10v (Error: %v)\n", res.URL, "ERR", res.Latency.Round(time.Millisecond), res.Err)
			continue
		}

		// Happy Path: Server responded successfully
		fmt.Printf("%-45s | %-6d | %-10v\n", res.URL, res.StatusCode, res.Latency.Round(time.Millisecond))
	}
}
