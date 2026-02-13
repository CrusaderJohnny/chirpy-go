package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/CrusaderJohnny/chirpy-go.git/internal/auth"
	"github.com/google/uuid"
)

type Webhook struct {
	Event string `json:"event"`
	Data  struct {
		UserID uuid.UUID `json:"user_id"`
	}
}

func (cfg *apiConfig) handlerWebHook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	webhook := Webhook{}
	err := decoder.Decode(&webhook)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could not decode request: ", err)
		return
	}
	if strings.ToLower(webhook.Event) != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "could not get api key from header: ", err)
		return
	}
	if apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "invalid api key: ", err)
		return
	}
	_, err = cfg.db.UpgradeToChirpyRed(r.Context(), webhook.Data.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "user not found: ", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "could not upgrade to chirpy red: ", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
