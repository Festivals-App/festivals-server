package handler

import (
	"database/sql"
	"github.com/Phisto/eventusserver/server/database"
	"github.com/Phisto/eventusserver/server/model"
	"net/http"
)

// GET functions

func GetArtists(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetObjects(db, "artist", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, artists)
}

func GetArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	artists, err := GetObject(db, "artist", objectID, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, artists)
}

func GetArtistImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	images, err := GetAssociatedImage(db, "artist", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func GetArtistLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	links, err := GetAssociatedLinks(db, "artist", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func GetArtistTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	tags, err := GetAssociatedTags(db, "artist", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tags)
}

func CreateArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artist, err := Create(db, r, "artist")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, artist)
}

func SetImageForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	resourceID, err := ResourceID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.SetResource(db, "artist", objectID, "image", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	resourceID, err := ResourceID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.SetResource(db, "artist", objectID, "link", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetTagForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	resourceID, err := ResourceID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.SetResource(db, "artist", objectID, "tag", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	resourceID, err := ResourceID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.RemoveResource(db, "artist", objectID, "image", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	resourceID, err := ResourceID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.RemoveResource(db, "artist", objectID, "link", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveTagForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	resourceID, err := ResourceID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.RemoveResource(db, "artist", objectID, "tag", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

// PATCH functions

func UpdateArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := Update(db, r, "artist")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, artists)
}

func DeleteArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.Delete(db, "artist", objectID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, []model.Artist{})
}
