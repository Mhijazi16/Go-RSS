package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with %v error %v\n", code, msg)
	}
	respondWithJson(w, code, struct {
		Error string `json:"error"`
	}{Error: msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload any) {
	var data, err = json.Marshal(payload)
	if err != nil {
		log.Fatalln("failed to parse json response!")
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
