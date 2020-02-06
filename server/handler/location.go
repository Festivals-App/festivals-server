package handler

import (
	"database/sql"
	"net/http"
)

func GetLocations(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetObjects(db, "location", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func GetLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetObject(db, r, "location")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func GetLocationImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "location", "image")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func GetLocationLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetAssociation(db, r, "location", "link")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func GetLocationPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetAssociation(db, r, "location", "place")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func SetImageForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "location", "image")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "location", "link")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetPlaceForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "location", "place")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "location", "image")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "location", "link")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemovePlaceForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "location", "place")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func CreateLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := Create(db, r, "location")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func UpdateLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := Update(db, r, "location")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func DeleteLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "location")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
