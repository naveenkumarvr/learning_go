/*
We have a new engineering problem to solve. Right now, our background workers are blindly shouting their results straight to the terminal screen using fmt.Printf. In a real production system, you don't want a chaotic mess of print statements. You want to aggregate those results into a final report, save them to a database, or trigger an alert system if a site goes down. Your first instinct might be to create a shared slice in main and have each goroutine append its results to it:

In Go, this is highly dangerous. A standard slice is not thread-safe. If two background workers finish at the exact same microsecond and try to write to the results slice at the exact same time, they will step on each other's toes, corrupt the memory, and crash your program with a fatal Data Race error.

"Do not communicate by sharing memory; instead, share memory by communicating."

Instead of sharing a vulnerable slice, Go gives us a thread-safe highway called a Channel (chan). Think of a channel like a secure pneumatic tube or a conveyor belt. One goroutine drops a message into the tube, and another goroutine safely catches it on the other side. No memory corruption, no overlapping writes, 100% thread-safe.

To Create
resultsChan := make(chan string)

To Send Data to channel
resultsChan <- "Success! Everything is valid."

To Read
msg := <-resultsChan

If you try to read channel and nothing comes in first the main program will wait for the message
*/

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func checkWebsite(url string, ch chan<- string) {
	resp, err := http.Get(url)

	// Gate1 : Check website connection if fails exit right away
	if err != nil {
		ch <- fmt.Sprintf("[DOWN] Connection error for %s: %v", url, err)
		return
	}

	// Gate2 : Close the http Connection once the operation is done. This will make sure connection is always closed at the end
	defer resp.Body.Close()

	// Gate3: Catch the status code error
	if resp.StatusCode != http.StatusOK {
		ch <- fmt.Sprintf("[ALERT] %s returned status %d", url, resp.StatusCode)
		return
	}
	ch <- fmt.Sprintf("[UP] %s is completely healthy", url)
}

func main() {
	targets := []string{
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/500",
		"https://this-domain-does-not-exist-at-all-12345.com",
	}

	//1. Initialize an unbuffered string channel
	resultChan := make(chan string)
	var wg sync.WaitGroup

	//2. Spin up workers
	for _, url := range targets {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			checkWebsite(u, resultChan)
		}(url)
	}

	for i := range targets {
		fmt.Print(i)

		fmt.Printf("%s\n", msg)
	}

	wg.Wait()
	fmt.Println("🏁 All concurrent checks completed successfully.")

}
