package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get chirps", err)
		return
	}
	formattedChirps := make([]Chirp, 0, len(chirps))
	authorString := r.URL.Query().Get("author_id")
	sortMethod := r.URL.Query().Get("sort")
	if authorString != "" {
		cleanedId, err := uuid.Parse(authorString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID: ", err)
			return
		}
		for _, c := range chirps {
			if c.UserID == cleanedId {
				formattedChirps = append(formattedChirps, Chirp{
					Id:        c.ID,
					CreatedAt: c.CreatedAt,
					UpdatedAt: c.UpdatedAt,
					Body:      c.Body,
					UserID:    c.UserID,
				})
			}
		}
	} else {
		for _, c := range chirps {
			formattedChirps = append(formattedChirps, Chirp{
				Id:        c.ID,
				CreatedAt: c.CreatedAt,
				UpdatedAt: c.UpdatedAt,
				Body:      c.Body,
				UserID:    c.UserID,
			})
		}
	}
	if sortMethod == "desc" {
		sort.Slice(formattedChirps, func(i, j int) bool {
			return formattedChirps[i].CreatedAt.After(formattedChirps[j].CreatedAt)
		})
	}
	respondWithJSON(w, http.StatusOK, formattedChirps)
}
