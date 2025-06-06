package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
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

func (db *apiConfig) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type requestBody struct {
		URL string `json:"url"`
	}

	var feedDTO requestBody
	var decoder = json.NewDecoder(r.Body)
	var err = decoder.Decode(&feedDTO)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to parse request body error: %s", err))
		return
	}

	feed, err := db.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Url:       feedDTO.URL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to create feed object, error: %s", err))
		return
	}

	respondWithJson(w, 201, toFeedDTO(feed))
}

func (db *apiConfig) getFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := db.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to fetch feed, error: %s", err))
	}

	respondWithJson(w, 200, toFeedsDTO(feeds))
}

func (db *apiConfig) FollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type requestBody struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	var body requestBody
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	follow, err := db.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    body.FeedId,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to follow feed, error: %s", err))
	}

	respondWithJson(w, 201, toFeedFollowDTO(follow))
}
