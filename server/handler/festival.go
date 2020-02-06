package handler

import (
	"database/sql"
	"net/http"
)

func GetFestivals(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := GetObjects(db, "festival", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, festivals)
}

func GetFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := GetObject(db, r, "festival")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, festivals)
}

func GetFestivalEvents(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := GetAssociation(db, r, "festival", "event")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, events)
}

func GetFestivalImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "festival", "image")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func GetFestivalLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetAssociation(db, r, "festival", "link")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func GetFestivalPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetAssociation(db, r, "festival", "place")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func GetFestivalTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetAssociation(db, r, "festival", "tag")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tags)
}

func SetEventForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "event")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetImageForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "image")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "link")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetPlaceForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "place")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetTagForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "tag")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "image")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "link")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemovePlaceForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "place")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveTagForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "tag")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func CreateFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festival, err := Create(db, r, "festival")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, festival)
}

func UpdateFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := Update(db, r, "festival")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, festivals)
}

func DeleteFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "festival")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
