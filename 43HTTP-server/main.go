package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Registering handlers
	/*
		type Handler interface {
			ServeHTTP(ResponseWriter, *Request)
		}
		A Handler responds to an HTTP request.
	*/

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)

	// Defining the server port
	port := "9080"
	log.Printf("Server starting on port %s...\n", port)

	// Starting the HTTP server
	err := http.ListenAndServe(":"+port, nil)
	/*
		// ListenAndServe always returns a non-nil error.
		func ListenAndServe(addr string, handler Handler) error {
			server := &Server{Addr: addr, Handler: handler}
			return server.ListenAndServe()
		}
	*/
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}

//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

// Handler for the root endpoint "/"
// A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Setting response header
	w.Header().Set("Content-Type", "text/plain")

	// Writing response
	fmt.Fprintln(w, "Welcome to our Go HTTP Server!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Creating a response object
	response := map[string]string{"message": "hello Mudasir "}

	// Encoding response to JSON and sending it
	w.Header().Set("content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}
