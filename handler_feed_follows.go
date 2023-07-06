package main

import (
	"encoding/json"
	"fmt"
	"github.com/MirjahonMirsaidov/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing the json, %v", err))
		return
	}
	feed_follow, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not create feed_follow, %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedFollowtoFeedFollow(feed_follow))
}

func (cfg apiConfig) handlerGetFeedFollowsList(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := cfg.DB.GetMyFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not retrieve feed follow list, %v", err))
		return
	}

	respondWithJson(w, 200, databaseFeedFollowstoFeedFollows(feedFollow))
}

func (cfg apiConfig) handlerUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not parse feed follow id, %v", err))
		return
	}
	err = cfg.DB.Unfollow(r.Context(), database.UnfollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not unfollow  this feed, %v", err))
		return
	}
	respondWithJson(w, 200, struct{}{})
}
