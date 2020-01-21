package database

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func Select(db *sql.DB, table string, objectID string) (*sql.Rows, error) {

	var query string
	var vars []interface{}
	// prepare select query
	if objectID == "" {
		query = "SELECT * FROM " + table + "s"
		vars = []interface{}{}
	} else {
		query = "SELECT * FROM " + table + "s WHERE " + table + "_id=?"
		vars = []interface{}{objectID}
	}
	// execute query
	return ExecuteRowQuery(db, query, vars)
}

func Search(db *sql.DB, table string, name string) (*sql.Rows, error) {

	var query string
	var vars []interface{}
	// prepare select query
	query = "SELECT * FROM " + table + "s WHERE " + table + "_name LIKE CONCAT('%', ?, '%')"
	vars = []interface{}{name}
	// execute query
	return ExecuteRowQuery(db, query, vars)
}

func Resource(db *sql.DB, object string, objectID string, resource string) (*sql.Rows, error) {

	var query string
	// prepare query
	if object == "event" {
		query = "SELECT * FROM " + resource + "s WHERE " + resource + "_id IN (SELECT `associated_" + resource + "` FROM events WHERE event_id=?)"
	} else if resource == "event" {
		query = "SELECT * FROM events WHERE `associated_" + object + "`=?"
	} else if object == "tag" {
		query = "SELECT * FROM " + resource + "s WHERE " + resource + "_id IN (SELECT `associated_" + resource + "` FROM `map_" + resource + "_" + object + "` WHERE `associated_" + object + "`=?)"
	} else {
		query = "SELECT * FROM " + resource + "s WHERE " + resource + "_id IN (SELECT `associated_" + resource + "` FROM `map_" + object + "_" + resource + "` WHERE `associated_" + object + "`=?)"
	}

	vars := []interface{}{objectID}
	// execute query
	return ExecuteRowQuery(db, query, vars)
}

func SetResource(db *sql.DB, table string, objectID string, resource string, resourceID string) error {

	if table == "event" {

		// prepare query
		query := "UPDATE `events` SET `associated_" + resource + "`=? WHERE `event_id`=?"
		vars := []interface{}{resourceID, objectID}
		// execute query
		result, err := ExecuteQuery(db, query, vars)
		if err != nil {
			return err
		}
		// check number of affected rows
		numOfAffectedRows, err := result.RowsAffected()
		if err != nil {
			return err
		}
		// we only delete one row per request
		if numOfAffectedRows != 1 {
			return errors.New("No rows where affected.")
		}
		return nil
	} else {

		// prepare query
		query := "INSERT INTO `map_" + table + "_" + resource + "` ( `associated_" + table + "` , `associated_" + resource + "` ) VALUES (?,?)"
		vars := []interface{}{objectID, resourceID}
		// execute query
		result, err := ExecuteQuery(db, query, vars)
		if err != nil {
			return err
		}
		_, err = result.LastInsertId()
		if err != nil {
			return err
		}
		return nil
	}
}

func Insert(db *sql.DB, table string, object interface{}) (*sql.Rows, error) {

	fields := DBFields(object)
	placeholder := DBPlaceholder(object)
	vars := DBValues(object)
	query := "INSERT INTO " + table + "s(" + fields + ") VALUES (" + placeholder + ")"

	result, err := ExecuteQuery(db, query, vars)
	if err != nil {
		return nil, err
	}

	instertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	objectID := strconv.FormatInt(instertID, 10)

	return Select(db, table, objectID)
}

func Update(db *sql.DB, table string, objectID string, object interface{}) (*sql.Rows, error) {

	keyValuePairs := DBKeyValuePairs(object)
	vars := DBValues(object)
	vars = append(vars, objectID) // for *table*_id value
	query := "UPDATE " + table + "s SET " + keyValuePairs + " WHERE `" + table + "_id`=?"

	log.Print(query)

	_, err := ExecuteQuery(db, query, vars)
	if err != nil {
		return nil, err
	}

	return Select(db, table, objectID)
}

func Delete(db *sql.DB, table string, objectID string) error {

	// prepare select query
	query := "DELETE FROM " + table + "s WHERE " + table + "_id=?"
	vars := []interface{}{objectID}

	log.Print(query)

	// execute query
	result, err := ExecuteQuery(db, query, vars)
	if err != nil {
		return err
	}
	// check number of affected rows
	numOfAffectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// we only delete one row per request
	if numOfAffectedRows != 1 {
		return errors.New("No rows where affected.")
	}
	return nil
}

/*
func associatedImage(objectName string, objectID string) (*sql.Rows, error) {

	queryString := "SELECT * FROM images WHERE image_id IN (SELECT `associated_image` FROM `map_" + objectName + "_image` WHERE `associated_" + objectName + "`=" + objectID + ")"
	return GetDB().Query(queryString)
}

func associatedLinks(objectName string, objectID string) (*sql.Rows, error) {

	queryString := "SELECT * FROM links WHERE link_id IN (SELECT `associated_link` FROM `map_" + objectName + "_link` WHERE `associated_" + objectName + "`=" + objectID + ")"
	return GetDB().Query(queryString)
}

func associatedPlace(objectName string, objectID string) (*sql.Rows, error) {

	queryString := "SELECT * FROM places WHERE place_id IN (SELECT `associated_place` FROM `map_" + objectName + "_place` WHERE `associated_" + objectName + "`=" + objectID + ")"
	return GetDB().Query(queryString)
}

func associatedTags(objectName string, objectID string) (*sql.Rows, error) {

	queryString := "SELECT * FROM tags WHERE tag_id IN (SELECT `associated_tag` FROM `map_" + objectName + "_tag` WHERE `associated_" + objectName + "`=" + objectID + ")"
	return GetDB().Query(queryString)
}

func associatedEvents(objectName string, objectID string) (*sql.Rows, error) {

	queryString := "SELECT * FROM events WHERE associated_" + objectName + "=" + objectID
	return GetDB().Query(queryString)
}
*/

// Execute a query that returns rows
func ExecuteRowQuery(db *sql.DB, query string, args []interface{}) (*sql.Rows, error) {

	rows, err := db.Query(query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return rows, nil
}

// Execute a query that does return a result
func ExecuteQuery(db *sql.DB, query string, args []interface{}) (sql.Result, error) {

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	stmt.Close()

	return result, nil
}
