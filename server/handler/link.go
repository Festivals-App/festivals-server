package handler

import (
	"database/sql"
	"net/http"

	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/rs/zerolog/log"
)

func GetLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetObjects(db, "link", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch links")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch links")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func GetLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetObject(db, r, "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch link")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch link")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func CreateLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := Create(db, r, "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to create link")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create link")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func UpdateLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := Update(db, r, "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to update link")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update link")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func DeleteLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete link")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete link")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
