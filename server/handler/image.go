package handler

import (
	"database/sql"
	"net/http"
)

func GetImages(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetObjects(db, "image", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func GetImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetObject(db, r, "image")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func CreateImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := Create(db, r, "image")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func UpdateImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := Update(db, r, "image")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func DeleteImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "image")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
