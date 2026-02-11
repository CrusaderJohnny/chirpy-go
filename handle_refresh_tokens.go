package main

import (
	"net/http"
	"time"

	"github.com/CrusaderJohnny/chirpy-go.git/internal/auth"
)

func (cfg *apiConfig) handleRefreshTokens(w http.ResponseWriter, r *http.Request) {
	type token struct {
		Token string `json:"token"`
	}
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't find token", err)
		return
	}
	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't find user", err)
		return
	}
	newJWToken, err := auth.MakeJWT(user.ID, cfg.secretToken, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't create new JWT token", err)
		return
	}
	respondWithJSON(w, http.StatusOK, token{
		Token: newJWToken,
	})
}
