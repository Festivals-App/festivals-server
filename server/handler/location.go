package handler

import (
	"database/sql"
	"net/http"

	"github.com/rs/zerolog/log"
)

func GetLocations(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetObjects(db, "location", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch locations")
		respondError(w, http.StatusBadRequest, "failed to fetch locations")
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func GetLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetObject(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch location")
		respondError(w, http.StatusBadRequest, "failed to fetch location")
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func GetLocationImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "location", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch image for location")
		respondError(w, http.StatusBadRequest, "failed to fetch image for location")
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func GetLocationLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetAssociation(db, r, "location", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch links for location")
		respondError(w, http.StatusBadRequest, "failed to fetch links for location")
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func GetLocationPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetAssociation(db, r, "location", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch place for location")
		respondError(w, http.StatusBadRequest, "failed to fetch place for location")
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func SetImageForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "location", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to set image for location")
		respondError(w, http.StatusBadRequest, "failed to set image for location")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "location", "link")
	if err != nil {
		log.Error().Err(err).Msg("")
		respondError(w, http.StatusBadRequest, "")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetPlaceForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "location", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to set place for location")
		respondError(w, http.StatusBadRequest, "failed to set place for location")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "location", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image of location")
		respondError(w, http.StatusBadRequest, "failed to remove image of location")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "location", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image from location")
		respondError(w, http.StatusBadRequest, "failed to remove image from location")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemovePlaceForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "location", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove place of location")
		respondError(w, http.StatusBadRequest, "failed to remove place of location")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func CreateLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := Create(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to create location")
		respondError(w, http.StatusBadRequest, "failed to create location")
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func UpdateLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := Update(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to update location")
		respondError(w, http.StatusBadRequest, "failed to update location")
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func DeleteLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete location")
		respondError(w, http.StatusBadRequest, "failed to delete location")
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
