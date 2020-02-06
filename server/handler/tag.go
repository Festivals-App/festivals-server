package handler

import (
	"database/sql"
	"net/http"
)

func GetTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetObjects(db, "tag", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tags)
}

func GetTag(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetObject(db, r, "tag")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tags)
}

func GetTagFestivals(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "tag", "festival")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func CreateTag(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := Create(db, r, "tag")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tags)
}

func UpdateTag(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := Update(db, r, "tag")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tags)
}

func DeleteTag(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "tag")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
