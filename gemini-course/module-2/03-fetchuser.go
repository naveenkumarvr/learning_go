package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Slideshow struct {
	Title  string `json: "title"`
	Author string `json: "author"`
}

type ApiResponse struct {
	Slideshow Slideshow `json: "slideshow"`
}

func fetchUserActivity() {

	// Adding Context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Gate 1
	req, err := http.NewReqeustWithContext(ctx, "GET", "https://httpbin.org/delay/1", nil)
	if err != nil {
		fmt.Printf("Error Creating Request: %v\n", err)
		return
	}
}
