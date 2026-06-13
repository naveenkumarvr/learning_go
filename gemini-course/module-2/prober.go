package main

import (
	"fmt"
	"io"
	"net/http"
)

func checkWebsite() {
	resp, err := http.Get("https://google.com")

	// Gate1: Check connection error (Here no connection is established)
	if err != nil {
		fmt.Printf("Connection error %v", err)
		return
	}

	// Gate2: Close the connection after running all the task
	// The reason why defer is called after Gate1 because at gate 1 the connection itself is not established
	defer resp.Body.Close()

	// Gate 3: Check Protocol Status (Here conneciton established)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Sever returned an error stauts: %v", resp.StatusCode)
		return
	}

	fmt.Println("Success! Everything is valid.")
	respBodyRaw, errRespBody := io.ReadAll(resp.Body)

	if errRespBody != nil {
		fmt.Printf("Error reading body %v", errRespBody)
		return
	}

	strResBody := string(respBodyRaw)
	fmt.Printf("Result:\n%s", strResBody)

}

func main() {

	checkWebsite()

}

/*
Topics Learned

1. Defer:
This tells go that run the specific piece of code at every exit you have from the function or at the end of the program. We can call it as schedule to run at the end

2. io.Readall:
This is a new method on from io package. Once the connection status code is 200 the server sends stream of data. Inorder to read all data go have special method called io.Readall. This reads all the stream of data and returns either err or respBody. The respBody will be of rawbyte. i.e byte slice []byte

3. Type conversion
In go we can convert the varible from one type to other. Syntax : TargetType(variable). eg: flaot64(integer)
*/
