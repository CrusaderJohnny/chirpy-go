package main

import (
	"encoding/json"
	"net/http"

	"github.com/CrusaderJohnny/chirpy-go.git/internal/auth"
	"github.com/CrusaderJohnny/chirpy-go.git/internal/database"
)

func (cfg *apiConfig) handleUserUpdate(w http.ResponseWriter, r *http.Request) {
	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "no access token found: ", err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "bad json format: ", err)
		return
	}
	defer r.Body.Close()
	if params.Password == "" || params.Email == "" {
		respondWithError(w, http.StatusBadRequest, "incorrect email or password format: ", err)
		return
	}
	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error hashing password: ", err)
		return
	}
	userId, err := auth.ValidateJWT(accessToken, cfg.secretToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "error getting user: ", err)
		return
	}
	updatedUser, err := cfg.db.UpdateUserById(r.Context(), database.UpdateUserByIdParams{
		ID:       userId,
		Email:    params.Email,
		Password: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error updating user: ", err)
		return
	}
	respondWithJSON(w, http.StatusOK, User{
		Id:        updatedUser.ID,
		Email:     updatedUser.Email,
		UpdatedAt: updatedUser.UpdatedAt,
		CreatedAt: updatedUser.CreatedAt,
	})
}
