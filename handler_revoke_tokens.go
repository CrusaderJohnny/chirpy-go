package main

import (
	"net/http"

	"github.com/CrusaderJohnny/chirpy-go.git/internal/auth"
)

func (cfg *apiConfig) handleRevokeTokens(w http.ResponseWriter, r *http.Request) {
	bearerToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "no token found", err)
		return
	}
	err = cfg.db.RevokeRefreshToken(r.Context(), bearerToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "could not revoke the specified token", err)
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}
