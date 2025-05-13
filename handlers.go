package main

import (
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{ Status string }{Status: "success"})
}

func handlAdmin(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "failed getting the admin panel!")
}
