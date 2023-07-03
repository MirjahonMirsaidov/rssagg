package main

import (
	"encoding/json"
	"fmt"
	"github.com/MirjahonMirsaidov/rssagg/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error parsing the json")
		return
	}
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not create user, %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedtoFeed(feed))
}

func (cfg apiConfig) handlerGetFeedList(w http.ResponseWriter, r *http.Request) {
	feed, err := cfg.DB.GetFeedList(r.Context())
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not retrieve feed list, %v", err))
	}

	respondWithJson(w, 200, databaseFeedstoFeeds(feed))
}
