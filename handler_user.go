package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rss/internal/auth"
	"rss/internal/database"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parammeters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parammeters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing json"))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("cannot create user"))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)

	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("invalid api key"))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("cannot get user"))
		return
	}
	

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}