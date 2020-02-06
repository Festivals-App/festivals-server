package handler

import (
	"database/sql"
	"net/http"
)

func GetPlaces(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetObjects(db, "place", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func GetPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetObject(db, r, "place")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func CreatePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := Create(db, r, "place")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func UpdatePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := Update(db, r, "place")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func DeletePlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "place")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
