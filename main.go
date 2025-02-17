package main

import (
	"fmt"
	"net/http"
)

// Handler function for the /hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with "Hello! Buddy"
	fmt.Fprintf(w, "Hello! Buddy")
}

func main() {
	// Set up the API route
	http.HandleFunc("/hello", helloHandler)

	// Start the HTTP server
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
