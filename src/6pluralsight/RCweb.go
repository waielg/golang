package main

import (
	"fmt"
	"net/http"
)

func healthcheckHANDLER(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)   //rc is 200
	fmt.Fprintln(w, "healthcheck") //response text
}

func main() {
	http.HandleFunc("/healthcheck", healthcheckHANDLER) // Register /healthcheck endpoint handler
	fmt.Println("Server is on :8080")                   // HTTP server on port 8080

	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

// To check HTTP status code, use the curl command in the terminal:
// curl -I http://localhost:8089/healthcheck
