package handler

import (
	"database/sql"
	"net/http"

	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/rs/zerolog/log"
)

func GetImages(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetObjects(db, "image", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch images")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch images")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func GetImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetObject(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch image")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch image")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func CreateImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := Create(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to create image")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create image")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func UpdateImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := Update(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to update image")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update image")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func DeleteImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete image")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete image")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
