package main

import (
	"fmt"
	"net/http"

	"github.com/mhijazi16/Go-RSS/auth"
	"github.com/mhijazi16/Go-RSS/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (db *apiConfig) authMiddleware(handler authHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("failed to authenticate user error: %s", err))
			return
		}

		user, err := db.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("failed to get user from database error: %s", err))
			return
		}

		handler(w, r, user)
	}
}
