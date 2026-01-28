package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	totalFailure := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.Handle("/", http.FileServer(http.Dir(".")))
	log.Println("Starting server on port 8080")
	err := totalFailure.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
