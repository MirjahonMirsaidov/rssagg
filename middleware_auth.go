package main

import (
	"fmt"
	"github.com/MirjahonMirsaidov/rssagg/internal/auth"
	"github.com/MirjahonMirsaidov/rssagg/internal/database"
	"net/http"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (cfg apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("%v", err))
			return
		}

		user, err := cfg.DB.GetUserByApikey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("User does not exists, %v", err))
			return
		}
		handler(w, r, user)
	}
}
