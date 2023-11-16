package handler

import (
	"database/sql"
	"net/http"

	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/rs/zerolog/log"
)

func GetArtists(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetObjects(db, "artist", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artists")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch artists")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, artists)
}

func GetArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetObject(db, r, "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, artists)
}

func GetArtistImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "artist", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artist image")
		servertools.RespondError(w, http.StatusInternalServerError, "failed to fetch artist image")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func GetArtistLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetAssociation(db, r, "artist", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artist links")
		servertools.RespondError(w, http.StatusInternalServerError, "failed to fetch artist links")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func GetArtistTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetAssociation(db, r, "artist", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artist tags")
		servertools.RespondError(w, http.StatusInternalServerError, "failed to fetch artist tags")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, tags)
}

func SetImageForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "artist", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to set image for artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set image for artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "artist", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to set link for artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set link for artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetTagForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "artist", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to set tag for artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set tag for artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "artist", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image from artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove image from artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "artist", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove link from artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove link from artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveTagForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "artist", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove tag from artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove tag from artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func CreateArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artist, err := Create(db, r, "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to create artist")
		servertools.RespondError(w, http.StatusInternalServerError, "failed to create artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, artist)
}

func UpdateArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := Update(db, r, "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to update artist")
		servertools.RespondError(w, http.StatusInternalServerError, "failed to update artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, artists)
}

func DeleteArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to delte artist")
		servertools.RespondError(w, http.StatusInternalServerError, "failed to delte artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
