package main

import (
	"time"

	"github.com/BowangLan/golang-rss-generator/internal/database"
)

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID.String(),
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
	}
}
