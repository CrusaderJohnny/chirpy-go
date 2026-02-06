package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	type returnError struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, returnError{msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}
