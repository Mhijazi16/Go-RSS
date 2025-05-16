package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/mhijazi16/Go-RSS/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func toUserDTO(user database.User) User {
	return User{
		Name:      user.Name,
		Password:  user.Password,
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
