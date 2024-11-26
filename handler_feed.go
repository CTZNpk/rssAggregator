package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/CTZNpk/rssaggregator/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithEror(w, 400, fmt.Sprintf("Error parinsg JSON: %v", err))
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      params.Name,
			Url:       params.URL,
			UserID:    user.ID,
		})

	if err != nil {
		respondWithEror(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedToFeed(feed))
}
