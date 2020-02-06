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

func GetObject(db *sql.DB, r *http.Request, entity string) ([]interface{}, error) {

	objectID, err := ObjectID(r)
	if err != nil {
		return nil, err
	}
	values := r.URL.Query()

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
		objcts, err := GetAssociatedObjects(db, entity, objectID, value, nil)
		if err != nil {
			return nil, err
		}
		relsDict[value] = objcts
	}
	return relsDict, nil
}

func GetAssociation(db *sql.DB, r *http.Request, entity string, association string) ([]interface{}, error) {

	objectID, err := ObjectID(r)
	if err != nil {
		return nil, err
	}
	includes := Includes(r)
	return GetAssociatedObjects(db, entity, objectID, association, includes)
}

func GetAssociatedObjects(db *sql.DB, entity string, objectID int, association string, includes []string) ([]interface{}, error) {

	rows, err := database.Resource(db, entity, objectID, association)
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
		obj, err := AnonScan(association, rows)
		if err != nil {
			return nil, err
		}
		if includes != nil {
			objId, err := AnonID(association, obj)
			if err != nil {
				return nil, err
			}
			includedRels, err := GetRelationships(db, association, objId, includes)
			if err != nil {
				return nil, err
			}
			obj, err = AnonInclude(association, obj, includedRels)
			if err != nil {
				return nil, err
			}
		}
		// add object result slice
		fetchedObjects = append(fetchedObjects, obj)
	}
	return fetchedObjects, nil
}

func SetAssociation(db *sql.DB, r *http.Request, entity string, association string) error {

	objectID, err := ObjectID(r)
	if err != nil {
		return err
	}
	resourceID, err := ResourceID(r)
	if err != nil {
		return err
	}
	err = database.SetResource(db, entity, objectID, association, resourceID)
	if err != nil {
		return err
	}
	return nil
}

func RemoveAssociation(db *sql.DB, r *http.Request, entity string, association string) error {

	objectID, err := ObjectID(r)
	if err != nil {
		return err
	}
	resourceID, err := ResourceID(r)
	if err != nil {
		return err
	}
	err = database.RemoveResource(db, entity, objectID, association, resourceID)
	if err != nil {
		return err
	}
	return nil
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

func Delete(db *sql.DB, r *http.Request, entity string) error {

	objectID, err := ObjectID(r)
	if err != nil {
		return err
	}
	err = database.Delete(db, entity, objectID)
	if err != nil {
		return err
	}
	return nil
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

func Includes(r *http.Request) []string {

	include := r.URL.Query().Get("include")
	if include != "" {
		includes, err := RelationshipNames(include)
		if err == nil {
			return includes
		}
	}
	return nil
}

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
