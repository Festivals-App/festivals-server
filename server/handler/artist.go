package handler

import (
	"database/sql"
	"net/http"
)

func GetArtists(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetObjects(db, "artist", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, artists)
}

func GetArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetObject(db, r, "artist")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, artists)
}

func GetArtistImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "artist", "image")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func GetArtistLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetAssociation(db, r, "artist", "link")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func GetArtistTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetAssociation(db, r, "artist", "tag")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tags)
}

func SetImageForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "artist", "image")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "artist", "link")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetTagForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "artist", "tag")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "artist", "image")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "artist", "link")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveTagForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "artist", "tag")
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func CreateArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artist, err := Create(db, r, "artist")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, artist)
}

func UpdateArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := Update(db, r, "artist")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, artists)
}

func DeleteArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "artist")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
