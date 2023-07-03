package main

import (
	"encoding/json"
	"fmt"
	"github.com/MirjahonMirsaidov/rssagg/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error parsing the json")
		return
	}
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not create user, %v", err))
		return
	}
	respondWithJson(w, 201, databaseUsertoUser(user))
}

func (cfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, 200, databaseUsertoUser(user))
}
