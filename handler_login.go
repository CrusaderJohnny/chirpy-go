package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/CrusaderJohnny/chirpy-go.git/internal/auth"
	"github.com/CrusaderJohnny/chirpy-go.git/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong", err)
		return
	}
	if params.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Password is required", errors.New("password is required"))
		return
	}
	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}
	correctPassword, err := auth.CheckHashPassword(params.Password, user.Password)
	if err != nil || !correctPassword {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}
	token, err := auth.MakeJWT(user.ID, cfg.secretToken, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Issue creating JWT: ", err)
		return
	}
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Issue creating Refresh Token: ", err)
		return
	}
	_, err = cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add((time.Hour * 24) * 60),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Issue saving refresh token to database: ", err)
		return
	}
	respondWithJSON(w, http.StatusOK, User{
		Id:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Email:        user.Email,
		Token:        token,
		RefreshToken: refreshToken,
		IsChirpyRed:  user.IsChirpyRed,
	})
}
