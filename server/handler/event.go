package handler

import (
	"database/sql"
	"net/http"
)

func GetEvents(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := GetObjects(db, "event", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, events)
}

func GetEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := GetObject(db, r, "event")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, events)
}

func GetEventFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := GetAssociation(db, r, "event", "festival")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, festivals)
}

func GetEventArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetAssociation(db, r, "event", "artist")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, artists)
}

func GetEventLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetAssociation(db, r, "event", "location")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func SetArtistForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "event", "artist")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetLocationForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "event", "location")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveArtistForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "event", "artist")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLocationForEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "event", "location")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func CreateEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := Create(db, r, "event")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, events)
}

func UpdateEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := Update(db, r, "event")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, events)
}

func DeleteEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "event")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
