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

func GetArtists(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var idValues []string
	// get query values if they exist
	values := r.URL.Query()
	if len(values) != 0 {

		// search with name
		name := values.Get("name")
		if name != "" {
			SearchArtists(name, db, w)
			return
		}
		// filter by ids
		ids := values.Get("ids")
		if ids != "" {
			var err error
			idValues, err = IDsFromString(ids)
			if err != nil {
				respondError(w, http.StatusBadRequest, err.Error())
				return
			}
		} else {
			respondError(w, http.StatusBadRequest, "Provided unknown query value")
			return
		}
	}

	if idValues == nil {
		idValues = []string{}
	}
	rows, err := database.Select(db, "artist", idValues)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Artist{})
	}
	var fetchedObjects []model.Artist
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.ArtistsScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

func SearchArtists(name string, db *sql.DB, w http.ResponseWriter) {

	rows, err := database.Search(db, "artist", name)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Artist{})
	}
	var fetchedObjects []model.Artist
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.ArtistsScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

func GetArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Select(db, "artist", []string{objectID})
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Artist{})
	}
	var fetchedObjects []model.Artist
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.ArtistsScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

func GetArtistImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Resource(db, "artist", objectID, "image")
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Image{})
	}
	var fetchedObjects []model.Image
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

func GetArtistLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Resource(db, "artist", objectID, "link")
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Link{})
	}
	var fetchedObjects []model.Link
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

func GetArtistTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Resource(db, "artist", objectID, "tag")
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Tag{})
	}
	var fetchedObjects []model.Tag
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

func CreateArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	body, readBodyErr := ioutil.ReadAll(r.Body)
	if readBodyErr != nil {
		respondError(w, http.StatusBadRequest, readBodyErr.Error())
		return
	}
	var objectToCreate model.Artist
	unmarshalErr := json.Unmarshal(body, &objectToCreate)
	if unmarshalErr != nil {
		respondError(w, http.StatusBadRequest, unmarshalErr.Error())
		return
	}
	if objectToCreate.Name == "" {
		respondError(w, http.StatusBadRequest, "The provided artist name is not valid.")
		return
	}
	rows, err := database.Insert(db, "artist", objectToCreate)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Artist{})
	}
	var fetchedObjects []model.Artist
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.ArtistsScan(rows)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	respondJSON(w, http.StatusOK, fetchedObjects)
}

func SetImageForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.SetResource(db, "artist", objectID, "image", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.SetResource(db, "artist", objectID, "link", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetTagForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.SetResource(db, "artist", objectID, "tag", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.RemoveResource(db, "artist", objectID, "image", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.RemoveResource(db, "artist", objectID, "link", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveTagForArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.RemoveResource(db, "artist", objectID, "tag", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

// PATCH functions

func UpdateArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	body, readBodyErr := ioutil.ReadAll(r.Body)
	if readBodyErr != nil {
		respondError(w, http.StatusBadRequest, readBodyErr.Error())
		return
	}
	var objectToUpdate model.Artist
	unmarshalErr := json.Unmarshal(body, &objectToUpdate)
	if unmarshalErr != nil {
		respondError(w, http.StatusBadRequest, unmarshalErr.Error())
		return
	}
	if objectToUpdate.Name == "" {
		respondError(w, http.StatusBadRequest, "The provided artist name is not valid.")
		return
	}
	objectID := chi.URLParam(r, "objectID")
	rows, err := database.Update(db, "artist", objectID, objectToUpdate)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Artist{})
	}
	var fetchedObjects []model.Artist
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.ArtistsScan(rows)
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

func DeleteArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	err := database.Delete(db, "artist", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, []model.Artist{})
}

/*
func GetArtistImage(w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := associatedImage("artist", objectID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var errors []Error
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

func GetArtistLinks(w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := associatedLinks("artist", objectID)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var fetchedLinks []Link
	for rows.Next() {

		// for each row, scan the result into our festival composite object
		link, err := LinksScan(rows)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		fetchedLinks = append(fetchedLinks, link)
	}

	render.JSON(w, r, fetchedLinks)
}

func GetArtistTags(w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	rows, err := associatedTags("artist", objectID)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var fetchedTags []Tag
	for rows.Next() {

		// for each row, scan the result into our festival composite object
		tag, err := TagsScan(rows)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		fetchedTags = append(fetchedTags, tag)
	}

	render.JSON(w, r, fetchedTags)
}

// POST functions

func CreateArtist(w http.ResponseWriter, r *http.Request) {

	response := make(map[string]string)
	response["message"] = "CreateArtist"
	render.JSON(w, r, response)
}

// PATCH functions

func UpdateArtistWithID(w http.ResponseWriter, r *http.Request) {

	response := make(map[string]string)
	response["message"] = "UpdateArtistWithID"
	render.JSON(w, r, response)
}

// DELETE functions

func DeleteArtistWithID(w http.ResponseWriter, r *http.Request) {

	response := make(map[string]string)
	response["message"] = "DeleteArtistWithID"
	render.JSON(w, r, response)
}
*/
