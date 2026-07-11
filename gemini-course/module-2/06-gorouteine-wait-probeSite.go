/*
Concurrency Race Condition: When your program starts, Go launches one primary thread called the Main Goroutine to run your main() function. When you put the word go in front of checkWebsite(url), the Main Goroutine does not wait for that network call to finish. Instead, it fires off a lightweight background worker and immediately rushes to the next line of code. Because network requests take time (hundreds of milliseconds) and launching background tasks takes almost no time (microseconds), the Main Goroutine finishes the loop, prints "🏁 All checks completed successfully.", and hits the closing curly brace of the main() function in a fraction of a millisecond.

The Golden Rule of Go Concurrency: The moment the main() function finishes, the entire program dies instantly. Go aggressively shuts down the application, killing all background workers before they even get a chance to finish dialing the servers or printing their results!

To fix this, we need a way to tell the Main Goroutine: "Hold on! Do not exit yet. You need to wait until all the background workers have finished their tasks before you close shop."

Go provides a built-in tool for this inside the standard library called a sync.WaitGroup. Think of a WaitGroup like a counter sheet at a construction site:

wg.Add(1): Increment the counter. You tell the manager, "I am sending 1 worker out to do a job."

wg.Done(): Decrement the counter. The worker yells back, "I am completely finished with my job!"

wg.Wait(): Block execution. The manager sits at the gate and refuses to let the program exit until the counter hits exactly 0.

To use a sync.WaitGroup, we need to import the "sync" package and update our main() loop.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
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

	// 1. Declare the WaitGroup tracker
	var wg sync.WaitGroup

	fmt.Println("Starting Sequential Health Check Engine")
	for _, url := range targets {

		// 2. Add 1 to the counter BEFORE spinning up the worker
		wg.Add(1)

		// 3. Spin up an independent background worker
		go func(u string) {
			// 4. Guarantee the counter drops when this wrapper function exits
			defer wg.Done()

			// 5. Run your airtight network client logic
			checkWebsite(u)
		}(url) //<-- We pass the current 'url' into the 'u' variable of the function
	}
	// 6. Freeze the main goroutine here until the counter hits exactly 0
	wg.Wait()
	fmt.Println("🏁 All checks completed successfully.")

}
