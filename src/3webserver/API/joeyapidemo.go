package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/mux"
)

// Declare a global variable to hold data
// Declare a mutex for concurrent access?
var (
	data     string
	dataLock sync.RWMutex
)

// Program execution starts here
func main() {
	// Create new instance of Gorilla Mux router
	router := mux.NewRouter()

	// Define routes and their corresponding handlers.
	router.HandleFunc("/health", HealthCheckHandler).Methods("GET") // Route to health check

	router.HandleFunc("/get", GetHandler).Methods("GET") // Route to retrieve data

	router.HandleFunc("/post", PostHandler).Methods("POST") // Route to recieve and store data via HTTP POST

	fmt.Println("Server is running on :8087")
	log.Fatal(http.ListenAndServe(":8087", router)) // Start HTTP server on port 8080 and handle requests using router
}

// Responds to GET requests on the /health route with ok
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // Set HTTP status code to 200 (OK)
	fmt.Fprintf(w, "OK")         // Write OK To response body
}

// Responds to GEAT requests for /get route by returning stored data
func GetHandler(w http.ResponseWriter, r *http.Request) {
	dataLock.RLock() // set a read lock for safety
	defer dataLock.RUnlock()

	w.WriteHeader(http.StatusOK)     // Set HTTP status code to 200
	fmt.Fprintf(w, "Data: %s", data) //Write the store data to the response body
}

// PostHandler responds to POST requests on the /post route by storing the received data
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	// if there is an error return an internal server error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// OTHERWISE put a write lock to safely update data
	dataLock.Lock()
	defer dataLock.Unlock()

	//Store the recieved data in the data variable
	data = string(body)
	// Set HTTP status code to 200
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data received and stored")
}

// curl -I http://localhost:8085/health
