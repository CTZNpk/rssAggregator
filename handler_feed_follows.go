package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/CTZNpk/rssaggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithEror(w, 400, fmt.Sprintf("Error parinsg JSON: %v", err))
	}

	feedFollow, err := apiConfig.DB.CreateFeedFollow(r.Context(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID:    params.FeedId,
			UserID:    user.ID,
		})

	if err != nil {
		respondWithEror(w, 400, fmt.Sprintf("Couldn't follow : %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiConfig *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithEror(w, 400, fmt.Sprintf("Couldn't get feed follows : %v", err))
		return
	}

	respondWithJson(w, 200, databseFeedFollowToFeedFollows(feedFollows))
}

func (apiConfig *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithEror(w, 400, fmt.Sprintf("Couldn't parse the feed follow id: %v", err))
		return
	}

	err = apiConfig.DB.DeleteFeedFollow(
		r.Context(),
		database.DeleteFeedFollowParams{
			ID:     feedFollowID,
			UserID: user.ID,
		})
	if err != nil {
		respondWithEror(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	respondWithJson(w, 200, struct{}{})
}
