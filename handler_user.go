package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/BowangLan/golang-rss-generator/internal/auth"
	"github.com/BowangLan/golang-rss-generator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Failed to parse JSON", err))
		return
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), userParams)
	if err != nil {
		respondWithError(w, 500, fmt.Sprint("Failed to create user", err))
		return
	}

	respondWithJson(w, 201, user)
}

func (apiCfg *apiConfig) handleListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.ListUsers(r.Context())
	if err != nil {
		respondWithError(w, 500, fmt.Sprint("Failed to list users", err))
		return
	}

	respondWithJson(w, 200, users)
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	if userID == "" {
		respondWithError(w, 400, "User ID is required")
		return
	}

	user, err := apiCfg.DB.GetUserById(r.Context(), uuid.MustParse(userID))
	if err != nil {
		respondWithError(w, 500, fmt.Sprint("Failed to get user", err))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetUserWithAuth(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKeyFromHeaders(r)
	if err != nil {
		respondWithError(w, 403, fmt.Sprint("Failed to get API key:", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 500, fmt.Sprint("Failed to get user", err))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))
}
