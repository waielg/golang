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
	http.HandleFunc("/healthcheck", healthcheckHANDLER) // register /healthcheck endpoint handler
	fmt.Println("Server is on :8080")                   // HTTp server on port
	http.ListenAndServe(":8080", nil)
}

// to check RC use curl command in terminal
// curl -I http://localhost:8080/healthcheck
