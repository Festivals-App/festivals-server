package handler

import (
	"database/sql"
	"net/http"
)

func GetLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetObjects(db, "link", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func GetLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetObject(db, r, "link")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func CreateLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := Create(db, r, "link")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func UpdateLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := Update(db, r, "link")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func DeleteLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "link")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
