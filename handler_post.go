package main

import (
	"fmt"
	"github.com/MirjahonMirsaidov/rssagg/internal/database"
	"net/http"
)

func (cfg apiConfig) handlerGetPostListForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.DB.GetPostListForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not retrieve feed list, %v", err))
	}

	respondWithJson(w, 200, databasePoststoPosts(posts))
}
