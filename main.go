package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("make sure you set PORT in .env file.")
	}

	fmt.Println(PORT)
}
