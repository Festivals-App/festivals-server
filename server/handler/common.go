package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/Phisto/eventusserver/server/database"
	"github.com/Phisto/eventusserver/server/model"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetObject(db *sql.DB, entity string, objectID int, values url.Values) ([]interface{}, error) {

	return GetObjects(db, entity, []int{objectID}, values)
}

func GetObjects(db *sql.DB, entity string, objectIDs []int, values url.Values) ([]interface{}, error) {

	var idValues []int
	var rels []string
	var err error
	idValues = append(idValues, objectIDs...)
	if len(values) != 0 {
		// search with name
		name := values.Get("name")
		if name != "" {
			return SearchObjects(db, entity, name)
		}
		// filter by ids
		ids := values.Get("ids")
		if ids != "" {
			var err error
			idValues, err = ObjectIDs(ids)
			if err != nil {
				return nil, err
			}
		}
		// handle include later
		include := values.Get("include")
		if include != "" {
			rels, err = RelationshipNames(include)
			if err != nil {
				return nil, err
			}
		}
	}
	rows, err := database.Select(db, entity, idValues)
	if err != nil {
		return nil, err
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		return nil, nil
	}
	var fetchedObjects []interface{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := AnonScan(entity, rows)
		if err != nil {
			return nil, err
		}
		if rels != nil {
			objId, err := AnonID(entity, obj)
			if err != nil {
				return nil, err
			}
			includedRels, err := GetRelationships(db, entity, objId, rels)
			if err != nil {
				return nil, err
			}
			obj, err = AnonInclude(entity, obj, includedRels)
			if err != nil {
				return nil, err
			}
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	return fetchedObjects, nil
}

func SearchObjects(db *sql.DB, entity string, name string) ([]interface{}, error) {

	rows, err := database.Search(db, entity, name)
	if err != nil {
		return nil, err
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		return []interface{}{}, nil
	}
	var fetchedObjects []interface{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := AnonScan(entity, rows)
		if err != nil {
			return nil, err
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	return fetchedObjects, nil
}

func GetRelationships(db *sql.DB, entity string, objectID int, relationships []string) (interface{}, error) {

	relsDict := make(map[string]interface{})
	for _, value := range relationships {
		if CompareSensitive(value, "event") {
			events, err := GetAssociatedEvents(db, entity, objectID, nil)
			if err != nil {
				return nil, err
			}
			relsDict[value] = events
		} else if CompareSensitive(value, "image") {
			images, err := GetAssociatedImage(db, entity, objectID)
			if err != nil {
				return nil, err
			}
			relsDict[value] = images
		} else if CompareSensitive(value, "links") {
			links, err := GetAssociatedLinks(db, entity, objectID)
			if err != nil {
				return nil, err
			}
			relsDict[value] = links
		} else if CompareSensitive(value, "place") {
			places, err := GetAssociatedPlace(db, entity, objectID)
			if err != nil {
				return nil, err
			}
			relsDict[value] = places
		} else if CompareSensitive(value, "tags") {
			tags, err := GetAssociatedTags(db, entity, objectID)
			if err != nil {
				return nil, err
			}
			relsDict[value] = tags
		} else if CompareSensitive(value, "festival") {
			festivals, err := GetAssociatedFestival(db, entity, objectID)
			if err != nil {
				return nil, err
			}
			relsDict[value] = festivals
		} else if CompareSensitive(value, "artist") {
			artists, err := GetAssociatedArtist(db, entity, objectID)
			if err != nil {
				return nil, err
			}
			relsDict[value] = artists
		} else if CompareSensitive(value, "location") {
			locations, err := GetAssociatedLocation(db, entity, objectID)
			if err != nil {
				return nil, err
			}
			relsDict[value] = locations
		} else {
			return nil, errors.New("get relationships: provided unknown relationship")
		}
	}
	return relsDict, nil
}

func GetAssociatedEvents(db *sql.DB, object string, objectID int, include []string) ([]model.Event, error) {

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
	if len(fetchedObjects) == 0 {
		fetchedObjects = []model.Event{}
	}
	return fetchedObjects, nil
}

func GetAssociatedImage(db *sql.DB, object string, objectID int) ([]model.Image, error) {

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
	if len(fetchedObjects) == 0 {
		fetchedObjects = []model.Image{}
	}
	return fetchedObjects, nil
}

func GetAssociatedLinks(db *sql.DB, object string, objectID int) ([]model.Link, error) {

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
	if len(fetchedObjects) == 0 {
		fetchedObjects = []model.Link{}
	}
	return fetchedObjects, nil
}

func GetAssociatedPlace(db *sql.DB, object string, objectID int) ([]model.Place, error) {

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
	if len(fetchedObjects) == 0 {
		fetchedObjects = []model.Place{}
	}
	return fetchedObjects, nil
}

func GetAssociatedTags(db *sql.DB, object string, objectID int) ([]model.Tag, error) {

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
	if len(fetchedObjects) == 0 {
		fetchedObjects = []model.Tag{}
	}
	return fetchedObjects, nil
}

func GetAssociatedFestival(db *sql.DB, object string, objectID int) ([]model.Festival, error) {

	rows, err := database.Resource(db, object, objectID, "festival")
	// check if an error occurred
	if err != nil {
		return nil, err
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		return []model.Festival{}, nil
	}
	var fetchedObjects []model.Festival
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.FestivalsScan(rows)
		if err != nil {
			return nil, err
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	if len(fetchedObjects) == 0 {
		fetchedObjects = []model.Festival{}
	}
	return fetchedObjects, nil
}

func GetAssociatedArtist(db *sql.DB, object string, objectID int) ([]model.Artist, error) {

	rows, err := database.Resource(db, object, objectID, "artist")
	// check if an error occurred
	if err != nil {
		return nil, err
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		return []model.Artist{}, nil
	}
	var fetchedObjects []model.Artist
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.ArtistsScan(rows)
		if err != nil {
			return nil, err
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	if len(fetchedObjects) == 0 {
		fetchedObjects = []model.Artist{}
	}
	return fetchedObjects, nil
}

func GetAssociatedLocation(db *sql.DB, object string, objectID int) ([]model.Location, error) {

	rows, err := database.Resource(db, object, objectID, "location")
	// check if an error occurred
	if err != nil {
		return nil, err
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		return []model.Location{}, nil
	}
	var fetchedObjects []model.Location
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := model.LocationsScan(rows)
		if err != nil {
			return nil, err
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	if len(fetchedObjects) == 0 {
		fetchedObjects = []model.Location{}
	}
	return fetchedObjects, nil
}

func Create(db *sql.DB, r *http.Request, entity string) ([]interface{}, error) {

	body, readBodyErr := ioutil.ReadAll(r.Body)
	if readBodyErr != nil {
		return nil, readBodyErr
	}
	objectToCreate, err := AnonUnmarshal(entity, body)
	if err != nil {
		return nil, err
	}
	rows, err := database.Insert(db, entity, objectToCreate)
	if err != nil {
		return nil, err
	}
	// no rows and no error indicate a successful query but an empty result
	if rows == nil {
		return []interface{}{}, nil
	}
	var fetchedObjects []interface{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := AnonScan(entity, rows)
		if err != nil {
			return nil, err
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	return fetchedObjects, nil
}

func Update(db *sql.DB, r *http.Request, entity string) ([]interface{}, error) {

	objectID, err := ObjectID(r)
	if err != nil {
		return nil, err
	}
	body, readBodyErr := ioutil.ReadAll(r.Body)
	if readBodyErr != nil {
		return nil, readBodyErr
	}
	objectToUpdate, err := AnonUnmarshal(entity, body)
	if err != nil {
		return nil, err
	}
	rows, err := database.Update(db, entity, objectID, objectToUpdate)
	if err != nil {
		return nil, err
	}
	var fetchedObjects []interface{}
	// iterate over the rows an create
	for rows.Next() {
		// scan the link
		obj, err := AnonScan(entity, rows)
		if err != nil {
			return nil, err
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	return fetchedObjects, nil
}

func ObjectID(r *http.Request) (int, error) {

	objectID := chi.URLParam(r, "objectID")
	num, err := strconv.ParseUint(objectID, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(num), nil
}

func ResourceID(r *http.Request) (int, error) {

	objectID := chi.URLParam(r, "resourceID")
	num, err := strconv.ParseUint(objectID, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(num), nil
}

// TODO Further string/value validation
func ObjectIDs(idsString string) ([]int, error) {

	var ids []int
	for _, id := range strings.Split(idsString, ",") {

		idNum, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, int(idNum))
	}

	if len(ids) == 0 {
		return nil, errors.New("ids parsing: failed to provide an id")
	}

	return ids, nil
}

func RelationshipNames(includes string) ([]string, error) {

	return strings.Split(includes, ","), nil
}

func AnonScan(entity string, rs *sql.Rows) (interface{}, error) {

	if CompareSensitive(entity, "event") {
		return model.EventsScan(rs)
	} else if CompareSensitive(entity, "festival") {
		return model.FestivalsScan(rs)
	} else if CompareSensitive(entity, "artist") {
		return model.ArtistsScan(rs)
	} else if CompareSensitive(entity, "location") {
		return model.LocationsScan(rs)
	} else if CompareSensitive(entity, "image") {
		return model.ImagesScan(rs)
	} else if CompareSensitive(entity, "link") {
		return model.LinksScan(rs)
	} else if CompareSensitive(entity, "place") {
		return model.PlacesScan(rs)
	} else if CompareSensitive(entity, "tag") {
		return model.TagsScan(rs)
	} else {
		return nil, errors.New("scan row: tried to scan an unknown entity")
	}
}

func AnonInclude(entity string, object interface{}, includes interface{}) (interface{}, error) {

	if CompareSensitive(entity, "event") {
		realObject := object.(model.Event)
		realObject.Include = includes
		return realObject, nil
	} else if CompareSensitive(entity, "festival") {
		realObject := object.(model.Festival)
		realObject.Include = includes
		return realObject, nil
	} else if CompareSensitive(entity, "artist") {
		realObject := object.(model.Artist)
		realObject.Include = includes
		return realObject, nil
	} else if CompareSensitive(entity, "location") {
		realObject := object.(model.Location)
		realObject.Include = includes
		return realObject, nil
	} else {
		return nil, errors.New("include relationship: tried to add relationships to an unknown entity")
	}
}

func AnonID(entity string, object interface{}) (int, error) {

	if CompareSensitive(entity, "event") {
		return object.(model.Event).ID, nil
	} else if CompareSensitive(entity, "festival") {
		return object.(model.Festival).ID, nil
	} else if CompareSensitive(entity, "artist") {
		return object.(model.Artist).ID, nil
	} else if CompareSensitive(entity, "location") {
		return object.(model.Location).ID, nil
	} else if CompareSensitive(entity, "image") {
		return object.(model.Image).ID, nil
	} else if CompareSensitive(entity, "link") {
		return object.(model.Link).ID, nil
	} else if CompareSensitive(entity, "place") {
		return object.(model.Place).ID, nil
	} else if CompareSensitive(entity, "tag") {
		return object.(model.Tag).ID, nil
	} else {
		return -1, errors.New("get id: tried to retrieve the ID of an unknown entity")
	}
}

func AnonUnmarshal(entity string, body []byte) (interface{}, error) {

	if CompareSensitive(entity, "event") {
		var objectToCreate model.Event
		err := json.Unmarshal(body, &objectToCreate)
		if err != nil {
			return nil, err
		}
		return objectToCreate, nil
	} else if CompareSensitive(entity, "festival") {
		var objectToCreate model.Festival
		err := json.Unmarshal(body, &objectToCreate)
		if err != nil {
			return nil, err
		}
		return objectToCreate, nil
	} else if CompareSensitive(entity, "artist") {
		var objectToCreate model.Artist
		err := json.Unmarshal(body, &objectToCreate)
		if err != nil {
			return nil, err
		}
		return objectToCreate, nil
	} else if CompareSensitive(entity, "location") {
		var objectToCreate model.Location
		err := json.Unmarshal(body, &objectToCreate)
		if err != nil {
			return nil, err
		}
		return objectToCreate, nil
	} else if CompareSensitive(entity, "image") {
		var objectToCreate model.Image
		err := json.Unmarshal(body, &objectToCreate)
		if err != nil {
			return nil, err
		}
		return objectToCreate, nil
	} else if CompareSensitive(entity, "link") {
		var objectToCreate model.Link
		err := json.Unmarshal(body, &objectToCreate)
		if err != nil {
			return nil, err
		}
		return objectToCreate, nil
	} else if CompareSensitive(entity, "place") {
		var objectToCreate model.Place
		err := json.Unmarshal(body, &objectToCreate)
		if err != nil {
			return nil, err
		}
		return objectToCreate, nil
	} else if CompareSensitive(entity, "tag") {
		var objectToCreate model.Tag
		err := json.Unmarshal(body, &objectToCreate)
		if err != nil {
			return nil, err
		}
		return objectToCreate, nil
	} else {
		return nil, errors.New("unmarshal object: tried to unmarshal an unknown entity")
	}
}

// taken from https://www.digitalocean.com/community/questions/how-to-efficiently-compare-strings-in-go
func CompareSensitive(a, b string) bool {
	// a quick optimization. If the two strings have a different
	// length then they certainly are not the same
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		// if the characters already match then we don't need to
		// alter their case. We can continue to the next rune
		if a[i] == b[i] {
			continue
		}
		if a[i] != b[i] {
			// the lowercase characters do not match so these
			// are considered a mismatch, break and return false
			return false
		}
	}
	// The string length has been traversed without a mismatch
	// therefore the two match
	return true
}
