package handler

import (
	"database/sql"
	"net/http"

	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/rs/zerolog/log"
)

func GetTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetObjects(db, "tag", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch tags")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch tags")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, tags)
}

func GetTag(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetObject(db, r, "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch tag")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch tag")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, tags)
}

func GetTagFestivals(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "tag", "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festivals for tag")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch festivals for tag")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func CreateTag(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := Create(db, r, "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to create tag")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create tag")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, tags)
}

func UpdateTag(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := Update(db, r, "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to update tag")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update tag")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, tags)
}

func DeleteTag(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete tag")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete tag")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
