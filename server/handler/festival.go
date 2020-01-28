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

	var idValues []string
	// get query values if they exist
	values := r.URL.Query()
	if len(values) != 0 {

		// search with name
		name := values.Get("name")
		if name != "" {
			SearchFestivals(name, db, w)
			return
		}
		// filter by ids
		ids := values.Get("ids")
		if ids != "" {
			var err error
			idValues, err = RelationshipsFromString(ids)
			if err != nil {
				respondError(w, http.StatusBadRequest, err.Error())
				return
			}
		} else {
			respondError(w, http.StatusBadRequest, "Provided unknown query value")
			return
		}
	}

	rows, err := database.Select(db, "festival", idValues)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Festival{})
	}
	var fetchedObjects []model.Festival
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

func SearchFestivals(name string, db *sql.DB, w http.ResponseWriter) {

	rows, err := database.Search(db, "festival", name)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Festival{})
	}
	var fetchedObjects []model.Festival
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
	rows, err := database.Select(db, "festival", []string{objectID})
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		respondJSON(w, http.StatusOK, []model.Festival{})
	}
	var fetchedObjects []model.Festival
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

	values := r.URL.Query()
	if len(values) != 0 {

		// get relationships to include
		includeQuery := values.Get("include")
		if includeQuery != "" {
			var err error
			stringVals, err := RelationshipsFromString(includeQuery)
			if err != nil {
				respondError(w, http.StatusBadRequest, err.Error())
				return
			}

			var includedRels []interface{}
			for _, value := range stringVals {
				if value == "events" {
					events, err := GetAssociatedEvents(db, "festival", objectID)
					// check if an error occurred
					if err != nil {
						respondError(w, http.StatusBadRequest, err.Error())
						return
					}
					resultMap := map[string]interface{}{value: events}
					includedRels = append(includedRels, resultMap)
				} else if value == "image" {
					images, err := GetAssociatedImage(db, "festival", objectID)
					// check if an error occurred
					if err != nil {
						respondError(w, http.StatusBadRequest, err.Error())
						return
					}
					resultMap := map[string]interface{}{value: images}
					includedRels = append(includedRels, resultMap)
				} else if value == "links" {
					links, err := GetAssociatedLinks(db, "festival", objectID)
					// check if an error occurred
					if err != nil {
						respondError(w, http.StatusBadRequest, err.Error())
						return
					}
					resultMap := map[string]interface{}{value: links}
					includedRels = append(includedRels, resultMap)
				} else if value == "place" {
					places, err := GetAssociatedPlace(db, "festival", objectID)
					// check if an error occurred
					if err != nil {
						respondError(w, http.StatusBadRequest, err.Error())
						return
					}
					resultMap := map[string]interface{}{value: places}
					includedRels = append(includedRels, resultMap)
				} else if value == "tags" {
					tags, err := GetAssociatedTags(db, "festival", objectID)
					// check if an error occurred
					if err != nil {
						respondError(w, http.StatusBadRequest, err.Error())
						return
					}
					resultMap := map[string]interface{}{value: tags}
					includedRels = append(includedRels, resultMap)
				} else {
					respondError(w, http.StatusBadRequest, "get festival: provided unknown relationship")
					return
				}
			}

			if len(includedRels) == 0 {
				respondError(w, http.StatusBadRequest, "get festival: provided unknown relationship")
				return
			}

			respondJSONWithIncludes(w, http.StatusOK, fetchedObjects, includedRels)
			return
		} else {
			respondError(w, http.StatusBadRequest, "get festival: provided unknown query value")
			return
		}
	} else {
		respondJSON(w, http.StatusOK, fetchedObjects)
	}
}

func GetFestivalEvents(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	events, err := GetAssociatedEvents(db, "festival", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, events)
}

func GetFestivalImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	images, err := GetAssociatedImage(db, "festival", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func GetFestivalLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	links, err := GetAssociatedLinks(db, "festival", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func GetFestivalPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	place, err := GetAssociatedPlace(db, "festival", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, place)
}

func GetFestivalTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	tags, err := GetAssociatedTags(db, "festival", objectID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tags)
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
	if rows == nil {
		respondError(w, http.StatusInternalServerError, "create festival: failed to insert into database")
		return
	}
	var fetchedObjects []model.Festival
	// iterate over the rows
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

func SetEventForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.SetResource(db, "festival", objectID, "event", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
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

func RemoveImageForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.RemoveResource(db, "festival", objectID, "image", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.RemoveResource(db, "festival", objectID, "link", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemovePlaceForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.RemoveResource(db, "festival", objectID, "place", resourceID)
	// check if an error occurred
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveTagForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	objectID := chi.URLParam(r, "objectID")
	resourceID := chi.URLParam(r, "resourceID")
	err := database.RemoveResource(db, "festival", objectID, "tag", resourceID)
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
	var fetchedObjects []model.Festival
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
