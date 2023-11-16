package handler

import (
	"database/sql"
	"net/http"

	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/rs/zerolog/log"
)

func GetEvents(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := GetObjects(db, "event", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch events")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch events")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, events)
}

func GetEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := GetObject(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, events)
}

func GetEventFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := GetAssociation(db, r, "event", "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch festival for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, festivals)
}

func GetEventImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetAssociation(db, r, "event", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch image for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch image for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, artists)
}

func GetEventArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetAssociation(db, r, "event", "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artist for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch artist for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, artists)
}

func GetEventLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetAssociation(db, r, "event", "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch location for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch location for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, locations)
}

func SetImageForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "event", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to set image for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set image for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetArtistForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "event", "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to set artist for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set artist for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetLocationForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "event", "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to set location for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set location for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "event", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image from event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove image from event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveArtistForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "event", "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove artist from event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove artist from event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLocationForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "event", "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove location from event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove location from event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func CreateEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := Create(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to create event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, events)
}

func UpdateEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := Update(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to update event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, events)
}

func DeleteEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
