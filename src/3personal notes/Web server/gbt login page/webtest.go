package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("./", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/unrecognized_path.html")
	})

	http.HandleFunc("/demopath1/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath1.html")
	})

	http.HandleFunc("/demopath1/subpatha", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopatha.html")
	})

	http.HandleFunc("/demopath2/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath2.html")
	})

	http.ListenAndServe(":8088", nil)

    err := http.ListenAndServe(":8088", nil)
    if err != nil {
        log.Fatal("HTTP server error: ", err)
    }
}

// url is localhost:8082/demopath2