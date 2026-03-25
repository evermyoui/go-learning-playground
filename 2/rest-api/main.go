package main

import (
	"log"
	"net/http"
)

func main() {

	port := ":8080"
	log.Printf("Starting server on localhost%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
