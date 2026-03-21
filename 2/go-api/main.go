package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(response, "Hello World")
}

func aboutHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "This is a Go APU - Stage 1!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about", aboutHandler)

	port := ":8080"
	log.Printf("Server starting on localhost: %v", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
