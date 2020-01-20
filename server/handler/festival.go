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

func GetFestivals(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	rows, err := database.Select(db, "festival", "")
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Festival{})
	}
	fetchedObjects := []model.Festival{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.FestivalsScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

func GetFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Select(db, "festival", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Festival{})
	}
	fetchedObjects := []model.Festival{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.FestivalsScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

func GetFestivalEvents(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	// prepare query
	query := "SELECT * FROM events WHERE associated_festival=?"
	vars := []interface{}{objectID}
	// execute query
	rows, err := database.ExecuteRowQuery(db, query, vars)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Event{})
	}
	fetchedObjects := []model.Event{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.EventsScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

func GetFestivalImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Resource(db, "festival", objectID, "image")
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Image{})
	}
	fetchedObjects := []model.Image{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.ImagesScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

func GetFestivalLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Resource(db, "festival", objectID, "link")
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

func GetFestivalPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Resource(db, "festival", objectID, "place")
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Place{})
	}
	fetchedObjects := []model.Place{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.PlacesScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

func GetFestivalTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Resource(db, "festival", objectID, "tag")
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Tag{})
	}
	fetchedObjects := []model.Tag{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.TagsScan(rows)
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

func CreateFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	body, readBodyErr := ioutil.ReadAll(r.Body)
	if readBodyErr != nil {
		respondError(w, http.StatusBadRequest, readBodyErr.Error())
		return
	}
	var objectToCreate model.Festival
	unmarshalErr := json.Unmarshal(body, &objectToCreate)
	if unmarshalErr != nil {
		respondError(w, http.StatusBadRequest, unmarshalErr.Error())
		return
	}
	if objectToCreate.Name == "" {
		respondError(w, http.StatusBadRequest, "The provided festival name is not valid.")
		return
	}
	rows, err := database.Insert(db, "festival", objectToCreate)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Festival{})
	}
	fetchedObjects := []model.Festival{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.FestivalsScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

func SetImageForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.SetResource(db, "festival", objectID, "image", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.SetResource(db, "festival", objectID, "link", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetPlaceForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.SetResource(db, "festival", objectID, "place", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetTagForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.SetResource(db, "festival", objectID, "tag", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

// PATCH functions

func UpdateFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	body, readBodyErr := ioutil.ReadAll(r.Body)
	if readBodyErr != nil {
		respondError(w, http.StatusBadRequest, readBodyErr.Error())
		return
	}
	var objectToUpdate model.Festival
	unmarshalErr := json.Unmarshal(body, &objectToUpdate)
	if unmarshalErr != nil {
		respondError(w, http.StatusBadRequest, unmarshalErr.Error())
		return
	}
	if objectToUpdate.Name == "" {
		respondError(w, http.StatusBadRequest, "The provided festival name is not valid.")
		return
	}
	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Update(db, "festival", objectID, objectToUpdate)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Festival{})
	}
	fetchedObjects := []model.Festival{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.FestivalsScan(rows)
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

func DeleteFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	err := database.Delete(db, "festival", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, []model.Festival{})
}

/*
func GetFestivalImage(w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := associatedImage("festival", objectID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var errors []model.Error
	var fetchedObjects []Image

	for rows.Next() {
		image, err := ImagesScan(rows)
		if err != nil {
			if err != sql.ErrNoRows {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else {
				fetchErr := Error{Code: 404, Detail: err.Error()}
				errors = append(errors, fetchErr)
			}
		}

		fetchedObjects = append(fetchedObjects, image)
	}

	RESPONSE(w, r, errors, fetchedObjects, nil)
}

func GetFestivalLinks(w http.ResponseWriter, r *http.Request) {


	objectID := chi.URLParam(r, "objectID")
	rows, err := associatedLinks("festival", objectID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var errors []model.Error
	var fetchedObjects []Link

	for rows.Next() {

		// for each row, scan the result into our festival composite object
		link, err := LinksScan(rows)
		if err != nil {
			if err != sql.ErrNoRows {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else {
				fetchErr := Error{Code: 404, Detail: err.Error()}
				errors = append(errors, fetchErr)
			}
		}

		fetchedObjects = append(fetchedObjects, link)
	}

	RESPONSE(w, r, errors, fetchedObjects, nil)
}

func GetFestivalPlace(w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := associatedPlace("festival", objectID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var errors []model.Error
	var fetchedObjects []Place

	for rows.Next() {

		place, err := PlacesScan(rows)
		if err != nil {
			if err != sql.ErrNoRows {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else {
				fetchErr := Error{Code: 404, Detail: err.Error()}
				errors = append(errors, fetchErr)
			}
		}

		fetchedObjects = append(fetchedObjects, place)
	}

	RESPONSE(w, r, errors, fetchedObjects, nil)
}

func GetFestivalTags(w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := associatedTags("festival", objectID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var errors []model.Error
	var fetchedObjects []Tag

	for rows.Next() {

		// for each row, scan the result into our festival composite object
		tag, err := TagsScan(rows)
		if err != nil {
			if err != sql.ErrNoRows {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else {
				fetchErr := Error{Code: 404, Detail: err.Error()}
				errors = append(errors, fetchErr)
			}
		}

		fetchedObjects = append(fetchedObjects, tag)
	}

	RESPONSE(w, r, errors, fetchedObjects, nil)
}

func GetFestivalEvents(w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := associatedEvents("festival", objectID)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var errors []model.Error
	var fetchedObjects []Event

	for rows.Next() {

		// for each row, scan the result into our festival composite object
		event, err := EventsScan(rows)
		if err != nil {
			if err != sql.ErrNoRows {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else {
				fetchErr := Error{Code: 404, Detail: err.Error()}
				errors = append(errors, fetchErr)
			}
		}

		fetchedObjects = append(fetchedObjects, event)
	}

	RESPONSE(w, r, errors, fetchedObjects, nil)
}
*/
