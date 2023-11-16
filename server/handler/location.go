package handler

import (
	"database/sql"
	"net/http"

	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/rs/zerolog/log"
)

func GetLocations(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetObjects(db, "location", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch locations")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch locations")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, locations)
}

func GetLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetObject(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, locations)
}

func GetLocationImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "location", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch image for location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch image for location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func GetLocationLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetAssociation(db, r, "location", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch links for location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch links for location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func GetLocationPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetAssociation(db, r, "location", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch place for location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch place for location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, places)
}

func SetImageForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "location", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to set image for location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set image for location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "location", "link")
	if err != nil {
		log.Error().Err(err).Msg("")
		servertools.RespondError(w, http.StatusBadRequest, "")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetPlaceForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "location", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to set place for location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set place for location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "location", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image of location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove image of location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "location", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image from location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove image from location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemovePlaceForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "location", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove place of location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove place of location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func CreateLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := Create(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to create location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, locations)
}

func UpdateLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := Update(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to update location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, locations)
}

func DeleteLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
