package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Phisto/eventusserver/server/database"
	"github.com/Phisto/eventusserver/server/model"
	"log"
	"net/http"
	"strings"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {

	//TODO String comparison is not very elegant!
	if fmt.Sprint(payload) == "[]" {
		payload = []interface{}{}
	}

	resultMap := map[string]interface{}{"data": payload}
	response, err := json.Marshal(resultMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Panic(err.Error())
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(response)
	if err != nil {
		log.Panic(err.Error())
	}
}

//TODO Refactor respondJSON with includes
// respondJSON makes the response with payload as json format
func respondJSONWithIncludes(w http.ResponseWriter, status int, payload interface{}, indlude interface{}) {

	//TODO String comparison is not very elegant!
	if fmt.Sprint(payload) == "[]" {
		payload = []interface{}{}
	}

	if fmt.Sprint(indlude) == "[]" {
		indlude = []interface{}{}
	}

	resultMap := map[string]interface{}{"data": payload, "indlude": indlude}
	response, err := json.Marshal(resultMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Panic(err.Error())
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(response)
	if err != nil {
		log.Panic(err.Error())
	}
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	resultMap := map[string]interface{}{"error": message}
	response, err := json.Marshal(resultMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Panic(err.Error())
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Panic(err.Error())
	}
}

// TODO Further string/value validation
func IDsFromString(ids string) ([]string, error) {

	return strings.Split(ids, ","), nil
}

func RelationshipsFromString(includes string) ([]string, error) {

	return strings.Split(includes, ","), nil
}

func GetAssociatedEvents(db *sql.DB, object string, objectID string) ([]model.Event, error) {

	rows, err := database.Resource(db, object, objectID, "event")
	// check if an error occurred
	if err != nil {
		return nil, err
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		return []model.Event{}, nil
	}
	var fetchedObjects []model.Event
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.EventsScan(rows)
		if err != nil {
			return nil, err
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	return fetchedObjects, nil
}

func GetAssociatedImage(db *sql.DB, object string, objectID string) ([]model.Image, error) {

	rows, err := database.Resource(db, object, objectID, "image")
	// check if an error occurred
	if err != nil {
		return nil, err
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		return []model.Image{}, nil
	}
	var fetchedObjects []model.Image
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.ImagesScan(rows)
		if err != nil {
			return nil, err
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	return fetchedObjects, nil
}

func GetAssociatedLinks(db *sql.DB, object string, objectID string) ([]model.Link, error) {

	rows, err := database.Resource(db, object, objectID, "link")
	// check if an error occurred
	if err != nil {
		return nil, err
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		return []model.Link{}, nil
	}
	var fetchedObjects []model.Link
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.LinksScan(rows)
		if err != nil {
			return nil, err
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	return fetchedObjects, nil
}

func GetAssociatedPlace(db *sql.DB, object string, objectID string) ([]model.Place, error) {

	rows, err := database.Resource(db, object, objectID, "place")
	// check if an error occurred
	if err != nil {
		return nil, err
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		return []model.Place{}, nil
	}
	var fetchedObjects []model.Place
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.PlacesScan(rows)
		if err != nil {
			return nil, err
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	return fetchedObjects, nil
}

func GetAssociatedTags(db *sql.DB, object string, objectID string) ([]model.Tag, error) {

	rows, err := database.Resource(db, object, objectID, "tag")
	// check if an error occurred
	if err != nil {
		return nil, err
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		return []model.Tag{}, nil
	}
	var fetchedObjects []model.Tag
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.TagsScan(rows)
		if err != nil {
			return nil, err
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	return fetchedObjects, nil
}
