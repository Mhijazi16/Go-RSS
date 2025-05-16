package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mhijazi16/Go-RSS/internal/database"
)

type apiConfig struct {
	DB *database.Queries
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

	DBURL := os.Getenv("DB_URL")
	if DBURL == "" {
		log.Fatal("make sure you set DBURL in .env file")
	}

	conn, err := sql.Open("postgres", DBURL)
	if err != nil {
		log.Fatal("failed to connect to database error: ", err)
	}

	apiCfg := apiConfig{DB: database.New(conn)}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	healthRouter := chi.NewRouter()
	healthRouter.Get("/health-check", handleHome)

	userRouter := chi.NewRouter()
	userRouter.Post("/", apiCfg.createUser)
	userRouter.Get("/", apiCfg.authMiddleware(apiCfg.getUser))

	router.Mount("/v1/users", userRouter)
	router.Mount("/v1", healthRouter)

	log.Printf("Starting server on port %s...\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
