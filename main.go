package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

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

	healthRouter := chi.NewRouter()
	healthRouter.Get("/health-check", handleHome)

	router.Mount("/v1", healthRouter)

	log.Printf("Starting server on port %s...\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
