package main

import (
	"net/http"

	"github.com/CrusaderJohnny/chirpy-go.git/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	type chirpReq struct {
		ID uuid.UUID `json:"id"`
	}
	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "no access token found: ", err)
		return
	}
	userID, err := auth.ValidateJWT(accessToken, cfg.secretToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "incorrect access token: ", err)
		return
	}
	idString := r.PathValue("chirpID")
	deleteId, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "no ID found: ", err)
		return
	}
	chirp, err := cfg.db.GetChirpById(r.Context(), deleteId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "no chirp found: ", err)
		return
	}
	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "incorrect user: ", err)
		return
	}
	err = cfg.db.DeleteChirpById(r.Context(), chirp.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not delete chirp: ", err)
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}
