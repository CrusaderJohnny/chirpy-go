package main

import "net/http"

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get chirps", err)
		return
	}
	formattedChirps := make([]Chirp, 0, len(chirps))
	for _, c := range chirps {
		formattedChirps = append(formattedChirps, Chirp{
			Id:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			Body:      c.Body,
			UserID:    c.UserID,
		})
	}
	respondWithJSON(w, http.StatusOK, formattedChirps)
}
