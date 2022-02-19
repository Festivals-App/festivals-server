package handler

import (
	"database/sql"
	"net/http"

	"github.com/rs/zerolog/log"
)

func GetImages(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetObjects(db, "image", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch images")
		respondError(w, http.StatusBadRequest, "failed to fetch images")
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func GetImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetObject(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch image")
		respondError(w, http.StatusBadRequest, "failed to fetch image")
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func CreateImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := Create(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to create image")
		respondError(w, http.StatusBadRequest, "failed to create image")
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func UpdateImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := Update(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to update image")
		respondError(w, http.StatusBadRequest, "failed to update image")
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func DeleteImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete image")
		respondError(w, http.StatusBadRequest, "failed to delete image")
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
