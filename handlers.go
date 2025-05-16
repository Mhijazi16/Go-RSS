package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mhijazi16/Go-RSS/auth"
	"github.com/mhijazi16/Go-RSS/internal/database"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{ Status string }{Status: "success"})
}

func (db *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	var userDTO parameters
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userDTO)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to parse json %s", err))
		return
	}

	user, err := db.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userDTO.Name,
		Password:  userDTO.Password,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to store object in database %s", err))
	}

	respondWithJson(w, 201, toUserDTO(user))
}

func (db *apiConfig) getUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, 200, toUserDTO(user))
}
