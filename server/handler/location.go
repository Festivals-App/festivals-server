package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/Phisto/eventusserver/server/database"
	"github.com/Phisto/eventusserver/server/model"
	"io/ioutil"
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

	body, readBodyErr := ioutil.ReadAll(r.Body)
	if readBodyErr != nil {
		respondError(w, http.StatusBadRequest, readBodyErr.Error())
		return
	}
	var objectToCreate model.Location
	unmarshalErr := json.Unmarshal(body, &objectToCreate)
	if unmarshalErr != nil {
		respondError(w, http.StatusBadRequest, unmarshalErr.Error())
		return
	}
	if objectToCreate.Name == "" {
		respondError(w, http.StatusBadRequest, "The provided location name is not valid.")
		return
	}
	rows, err := database.Insert(db, "location", objectToCreate)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Location{})
	}
	var fetchedObjects []model.Location
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.LocationsScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
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

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	body, readBodyErr := ioutil.ReadAll(r.Body)
	if readBodyErr != nil {
		respondError(w, http.StatusBadRequest, readBodyErr.Error())
		return
	}
	var objectToUpdate model.Location
	unmarshalErr := json.Unmarshal(body, &objectToUpdate)
	if unmarshalErr != nil {
		respondError(w, http.StatusBadRequest, unmarshalErr.Error())
		return
	}
	if objectToUpdate.Name == "" {
		respondError(w, http.StatusBadRequest, "The provided location name is not valid.")
		return
	}
	rows, err := database.Update(db, "location", objectID, objectToUpdate)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Location{})
	}
	var fetchedObjects []model.Location
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.LocationsScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

// DELETE functions

func DeleteLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID, err := ObjectID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.Delete(db, "location", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, []model.Artist{})
}
