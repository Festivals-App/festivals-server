package handler

import (
	"database/sql"
	"github.com/Phisto/eventusserver/server/database"
	"github.com/Phisto/eventusserver/server/model"
	"net/http"
)

// GET functions

func GetPlaces(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetObjects(db, "place", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func GetPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	places, err := GetObject(db, "place", objectID, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, places)
}

// POST functions

func CreatePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := Create(db, r, "place")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, places)
}

// PATCH functions

func UpdatePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := Update(db, r, "place")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, places)
}

// DELETE functions

func DeletePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.Delete(db, "place", objectID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, []model.Place{})
}
