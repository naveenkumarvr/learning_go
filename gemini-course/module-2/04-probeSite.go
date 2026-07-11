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
		checkWebsite(url)
	}
	fmt.Println("🏁 All checks completed successfully.")

}
