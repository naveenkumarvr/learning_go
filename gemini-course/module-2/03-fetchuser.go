package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Slideshow struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type ApiResponse struct {
	Slideshow Slideshow `json:"slideshow"`
}

func fetchUserActivity() {

	// Adding Context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Gate 1
	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/json", nil)
	if err != nil {
		fmt.Printf("Error Creating Request: %v\n", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Connection error %v", err)
		return
	}
	defer resp.Body.Close()
	var apiResp ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		fmt.Printf("Error decoding data \n %v", err)
		return
	}
	fmt.Printf("Decoding Successful")
	fmt.Printf("%v\n", apiResp)
	fmt.Printf("Title: %s\n", apiResp.Slideshow.Title)
	fmt.Printf("Author: %s\n", apiResp.Slideshow.Author)

}

func main() {
	fetchUserActivity()
}
