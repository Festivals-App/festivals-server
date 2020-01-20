package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/Phisto/eventusserver/server/database"
	"github.com/Phisto/eventusserver/server/model"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

// GET functions

func GetLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	rows, err := database.Select(db, "link", "")
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Link{})
	}
	fetchedObjects := []model.Link{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.LinksScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

func GetLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Select(db, "link", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Link{})
	}
	fetchedObjects := []model.Link{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.LinksScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

// POST functions

func CreateLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	body, readBodyErr := ioutil.ReadAll(r.Body)
	if readBodyErr != nil {
		respondError(w, http.StatusBadRequest, readBodyErr.Error())
		return
	}
	var objectToCreate model.Link
	unmarshalErr := json.Unmarshal(body, &objectToCreate)
	if unmarshalErr != nil {
		respondError(w, http.StatusBadRequest, unmarshalErr.Error())
		return
	}
	rows, err := database.Insert(db, "link", objectToCreate)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Link{})
	}
	fetchedObjects := []model.Link{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.LinksScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

// PATCH functions

func UpdateLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	body, readBodyErr := ioutil.ReadAll(r.Body)
	if readBodyErr != nil {
		respondError(w, http.StatusBadRequest, readBodyErr.Error())
		return
	}
	var objectToUpdate model.Link
	unmarshalErr := json.Unmarshal(body, &objectToUpdate)
	if unmarshalErr != nil {
		respondError(w, http.StatusBadRequest, unmarshalErr.Error())
		return
	}
	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Update(db, "link", objectID, objectToUpdate)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Link{})
	}
	fetchedObjects := []model.Link{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.LinksScan(rows)
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

func DeleteLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	err := database.Delete(db, "link", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, []model.Link{})
}
