package handler

import (
	"database/sql"
	"github.com/Phisto/eventusserver/server/database"
	"github.com/Phisto/eventusserver/server/model"
	"net/http"
)

// GET functions

func GetLocations(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetObjects(db, "location", nil, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func GetLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	locations, err := GetObject(db, "location", objectID, r.URL.Query())
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func GetLocationImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	images, err := GetAssociatedImage(db, "location", objectID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func GetLocationLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	links, err := GetAssociatedLinks(db, "location", objectID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func GetLocationPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	places, err := GetAssociatedPlace(db, "location", objectID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, places)
}

// POST functions

func CreateLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := Create(db, r, "location")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

func SetImageForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

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
	err = database.SetResource(db, "location", objectID, "image", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

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
	err = database.SetResource(db, "location", objectID, "link", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetPlaceForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

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
	err = database.SetResource(db, "location", objectID, "place", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

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
	err = database.RemoveResource(db, "location", objectID, "image", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

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
	err = database.RemoveResource(db, "location", objectID, "link", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemovePlaceForLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

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
	err = database.RemoveResource(db, "location", objectID, "place", resourceID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

// PATCH functions

func UpdateLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := Update(db, r, "location")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, locations)
}

// DELETE functions

func DeleteLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.Delete(db, "location", objectID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, []model.Artist{})
}
