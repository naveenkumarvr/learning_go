package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates Gin Router instance
	r := gin.Default()
	// A rounter Engine is created. Logging Middleware(Log Requests). Recovery Middleware(recovers from panics and avoids server crash)

	r.GET("/ping", func(c *gin.Context) {
		/*
			r.Get : Register a handler for the HTTP Get method
			"/ping": Route path (localhost:8080/ping)
			"func(c *gin.Context)": This handler function for this route
				c *gin.Context gives access to the request, response, query parameters, etc.
		*/
		c.JSON(http.StatusOK, gin.H{
			// This line sends response to the received request with the stauts HTTP OK(200).
			// gin.H is a helper (a map of alias that returns)
			"message": "pong",
		})
	})
	r.Run()
	// This starts the HTTP server on port 8080 by default.
	// It listen on 0.0.0.0:8080. we can modify port by running r.Run(":3000") or we can update ip by using r.Run("127.0.0.1:3000")
}
