package main

import (
	"errors"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Platform is only available for admins", errors.New("not authorized"))
		return
	}
	cfg.fileserverHits.Store(0)
	err := cfg.db.Reset(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to reset the database", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0 and database reset to initial state"))
}
