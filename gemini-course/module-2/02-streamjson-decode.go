package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserActivity struct {
	Username      string `json:"username"`
	LoginAttempts int    `json:"login_attempts"`
}

func validateWebsite(website string) bool {
	resp, err := http.Get(website)

	// Gate 1: Conneciton Error
	if err != nil {
		fmt.Printf("Connection error : %v", err)
		return false
	}

	// Gate 2 Closing the open connection
	defer resp.Body.Close()

	//Gate3
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Sever returned an error status: %v", resp.StatusCode)
		return false
	}
	fmt.Printf("Success! Everything is valid")
	return true

}

func main() {
	websiteValidation := validateWebsite("https://google.com")
	if websiteValidation {
		var activity UserActivity

		//json.NewDecoder(resp.Body) creates a parsing tool that hooks directly up to the active network pipe, and .Decode(&activity) streams the data straight into your struct variables.
		if err := json.NewDecoder(resp.Body).Decode(&activity); err != nil {
			fmt.Printf("Error decoding data.\n%v", err)
			return
		}
		fmt.Printf("Decoding successful")

		fmt.Printf("\nUsername: %s, LoginAttempts: %d", activity.Username, activity.LoginAttempts)
	}
}

/*
Struct Variable Capitalization: We can see that "Username" Variable is capitalized which is necessary in go to indicate that the var is public, and can and will be accessed by any package. If it is small case it is private to specific package. In our example since this variable is accessed by enconding/json which is public lib we make it as Capital. If we make it small the lib just ignores this variable

JsonDecoder: we used json.Unmarshal, which takes a complete slice of bytes already stored in memory. With network connections, however, the data arrives as a live stream (io.Reader).
Instead of reading the entire network stream into memory first using io.ReadAll and then unmarshalling it, Go gives us a tool to decode the data directly from the wire: json.NewDecoder.

*/
