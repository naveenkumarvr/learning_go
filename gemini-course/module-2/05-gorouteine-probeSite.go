/*
Imagine you are monitoring 100 enterprise servers, and a few of them are running slow, taking 2 seconds each to respond. If you check them one by one, your script will take over 3 or 4 minutes just to complete a single loop! In production monitoring, that latency is unacceptable. We want to check all 100 servers at the exact same split-second, so the entire program takes only 2 seconds total to complete.

In Go, we do this by spinning up a Goroutine. A goroutine is a lightweight thread managed by the Go runtime.

To turn any standard function call into an independent background task that runs concurrently, you simply put the go keyword right in front of the function call:

Concurrency Race Condition: When your program starts, Go launches one primary thread called the Main Goroutine to run your main() function. When you put the word go in front of checkWebsite(url), the Main Goroutine does not wait for that network call to finish. Instead, it fires off a lightweight background worker and immediately rushes to the next line of code. Because network requests take time (hundreds of milliseconds) and launching background tasks takes almost no time (microseconds), the Main Goroutine finishes the loop, prints "🏁 All checks completed successfully.", and hits the closing curly brace of the main() function in a fraction of a millisecond.

The Golden Rule of Go Concurrency: The moment the main() function finishes, the entire program dies instantly. Go aggressively shuts down the application, killing all background workers before they even get a chance to finish dialing the servers or printing their results!
*/

package main

import (
	"fmt"
	"io"
	"net/http"
)

func checkWebsite(url string) {
	resp, err := http.Get(url)

	// Gate1: Check connection error (Here no connection is established)
	if err != nil {
		fmt.Printf("Connection error %v\n", err)
		return
	}

	// Gate2: Close the connection after running all the task
	// The reason why defer is called after Gate1 because at gate 1 the connection itself is not established
	defer resp.Body.Close()

	// Gate 3: Check Protocol Status (Here conneciton established)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Sever returned an error stauts: %v\n", resp.StatusCode)
		return
	}

	fmt.Println("Success! Everything is valid.")
	respBodyRaw, errRespBody := io.ReadAll(resp.Body)

	if errRespBody != nil {
		fmt.Printf("Error reading body %v\n", errRespBody)
		return
	}

	strResBody := string(respBodyRaw)
	fmt.Printf("Result:\n%s", strResBody)

}

func main() {

	targets := []string{
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/500",
		"https://this-domain-does-not-exist-at-all-12345.com",
	}

	fmt.Println("Starting Sequential Health Check Engine")
	for _, url := range targets {
		// 🚀 Concurrent: Main fires this off into the background and instantly moves to the next line
		go checkWebsite(url)
	}
	fmt.Println("🏁 All checks completed successfully.")

}
