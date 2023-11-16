package handler

import (
	"database/sql"
	"net/http"

	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/rs/zerolog/log"
)

func GetPlaces(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetObjects(db, "place", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch places")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch places")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, places)
}

func GetPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetObject(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch place")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch place")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, places)
}

func CreatePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := Create(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to create place")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create place")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, places)
}

func UpdatePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := Update(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to update place")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update place")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, places)
}

func DeletePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete place")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete place")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
