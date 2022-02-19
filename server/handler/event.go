package handler

import (
	"database/sql"
	"net/http"

	"github.com/rs/zerolog/log"
)

func GetEvents(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := GetObjects(db, "event", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch events")
		respondError(w, http.StatusBadRequest, "failed to fetch events")
		return
	}
	respondJSON(w, http.StatusOK, events)
}

func GetEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := GetObject(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch event")
		respondError(w, http.StatusBadRequest, "failed to fetch event")
		return
	}
	respondJSON(w, http.StatusOK, events)
}

func GetEventFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := GetAssociation(db, r, "event", "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival for event")
		respondError(w, http.StatusBadRequest, "failed to fetch festival for event")
		return
	}
	respondJSON(w, http.StatusOK, festivals)
}

func GetEventImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetAssociation(db, r, "event", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch image for event")
		respondError(w, http.StatusBadRequest, "failed to fetch image for event")
		return
	}
	respondJSON(w, http.StatusOK, artists)
}

func GetEventArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetAssociation(db, r, "event", "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artist for event")
		respondError(w, http.StatusBadRequest, "failed to fetch artist for event")
		return
	}
	respondJSON(w, http.StatusOK, artists)
}

func GetEventLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetAssociation(db, r, "event", "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch location for event")
		respondError(w, http.StatusBadRequest, "failed to fetch location for event")
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func SetImageForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "event", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to set image for event")
		respondError(w, http.StatusBadRequest, "failed to set image for event")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetArtistForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "event", "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to set artist for event")
		respondError(w, http.StatusBadRequest, "failed to set artist for event")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetLocationForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "event", "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to set location for event")
		respondError(w, http.StatusBadRequest, "failed to set location for event")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "event", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image from event")
		respondError(w, http.StatusBadRequest, "failed to remove image from event")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveArtistForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "event", "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove artist from event")
		respondError(w, http.StatusBadRequest, "failed to remove artist from event")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLocationForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "event", "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove location from event")
		respondError(w, http.StatusBadRequest, "failed to remove location from event")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func CreateEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := Create(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to create event")
		respondError(w, http.StatusBadRequest, "failed to create event")
		return
	}
	respondJSON(w, http.StatusOK, events)
}

func UpdateEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := Update(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to update event")
		respondError(w, http.StatusBadRequest, "failed to update event")
		return
	}
	respondJSON(w, http.StatusOK, events)
}

func DeleteEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete event")
		respondError(w, http.StatusBadRequest, "failed to delete event")
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
