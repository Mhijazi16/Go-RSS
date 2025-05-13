package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type response struct {
	Status string
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var data = response{Status: "success"}
	var res, err = json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(res)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("make sure you set PORT in .env file.")
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/health-check", handleHome)

	log.Printf("Starting server on port %s...\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
