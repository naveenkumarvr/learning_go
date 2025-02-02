package main

import (
	"fmt"
	"log"
	"net/http"
)

// Hello Handler contains two things the response which is referred as parameter w and the request from the http which is referred as request r. Here the request r is referred in pointer because we have to Point to the http request sent by the user.
// http.ResponseWriter handles sending response to the request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// We are handling error if user types a wrong word than hello here. r is the request received which contains many things in which we are using URL parameter and inside that we are checking the path
	if r.URL.Path != "/hello" {
		// If not /hello then we handle error with http.Error which accepts 3 inputs by default, The response, and string notification for user and http response code
		http.Error(w, "404 Not Found", http.StatusNotFound)
		// If error occurs return from the function itself
		return
	}

	//Handling Error for method. If users try to send the request other than GET we should send error
	// Again from the request pointer r we are checking method we get
	if r.Method != "GET" {
		http.Error(w, "method is not Supported. Only Get method is supported", http.StatusNotFound)
		return
	}
	// Finally if no error return hello as response
	//Fprintf formats according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered.
	fmt.Fprintf(w, "Hello!!!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// We have to return error if the http form which is submitted not correct. Or unable to parse it. Parse is nothing but reading the form data
	// The parse form results are captured in err var and if it returns error that is handled in other section
	if err := r.ParseForm(); err != nil {
		// using formatted print to return the response in case of error with the formatter %v
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if r.FormValue("name") == "" || r.FormValue("address") == "" {
		http.Error(w, "Error: Form value is empty", http.StatusUnprocessableEntity)
		return
	}
	// Finally returning the response
	fmt.Fprintf(w, "POST request successful\n")
	// Trying to return the Captured value for our reference. We are extracting the data from FormValue
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name:%s\n", name)
	fmt.Fprintf(w, "Address:%s\n", address)
}

func main() {
	// created a variable called fileserver which create an Fileserver Handler which will serve to the http request
	// We refer the FileServer handler to get the static content from the static folder. By Default it will pick the index.html as default page
	fileServer := http.FileServer(http.Dir("./static"))

	// http Handle to handle the patterns path in the webserver. In our case we are saying the root request should be redirected to our fileServer variable which have the FileServer which will serve the http request
	http.Handle("/", fileServer)

	// Now if we want to handle a Pattern path with a custom function such as do X when user hit this path we have to use HandleFunc of http
	// In the below line we are handling hello and form pattern with our own custom function.
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting the webserver at port 8080\n")

	// http.ListenAndServe Function creates the http server which handles the request. This is the core of webserver80
	// In the below code we ask the ListenAndServe to listen on port 8080.
	// the nil at the ListenAndServe refers to where we are routing the requests, in our case we say default which is the http will look for http.Handle and route it to there. If you have created a separate handler you can add that here. Since we are asking it to route to default Handler i.e http.Handle we are using nil
	// This function returns two values response and error status.
	// In the next section we are handling error where if the port is used or handler is not responding the err will not able to reach http.Handle we will get error that is handled
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
