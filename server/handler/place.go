package handler

import (
	"database/sql"
	"net/http"

	"github.com/rs/zerolog/log"
)

func GetPlaces(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetObjects(db, "place", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch places")
		respondError(w, http.StatusBadRequest, "failed to fetch places")
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func GetPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetObject(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch place")
		respondError(w, http.StatusBadRequest, "failed to fetch place")
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func CreatePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := Create(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to create place")
		respondError(w, http.StatusBadRequest, "failed to create place")
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func UpdatePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := Update(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to update place")
		respondError(w, http.StatusBadRequest, "failed to update place")
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func DeletePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete place")
		respondError(w, http.StatusBadRequest, "failed to delete place")
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
