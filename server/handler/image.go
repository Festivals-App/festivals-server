package handler

import (
	"database/sql"
	"github.com/Phisto/eventusserver/server/database"
	"github.com/Phisto/eventusserver/server/model"
	"net/http"
)

// GET functions

func GetImages(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetObjects(db, "image", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func GetImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	images, err := GetObject(db, "image", objectID, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

// POST functions

func CreateImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := Create(db, r, "image")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

// PATCH functions

func UpdateImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := Update(db, r, "image")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

// DELETE functions

func DeleteImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.Delete(db, "image", objectID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, []model.Image{})
}
