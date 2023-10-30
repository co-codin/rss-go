package main

import (
	"fmt"
	"net/http"
	"rss/internal/auth"
	"rss/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("invalid api key"))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("cannot get user"))
			return
		}

		handler(w, r, user)
	}
}
