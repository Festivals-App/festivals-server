package database

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func Select(db *sql.DB, table string, objectIDs []string) (*sql.Rows, error) {

	var query string
	var vars []interface{}
	// prepare select query
	if len(objectIDs) == 0 {
		query = "SELECT * FROM " + table + "s;"
		vars = []interface{}{}
	} else {

		if len(objectIDs) == 1 {
			query = "SELECT * FROM " + table + "s WHERE " + table + "_id=?;"
			vars = []interface{}{objectIDs[0]}
		} else {
			placeholder := DBPlaceholderForIDs(objectIDs)
			query = "SELECT * FROM " + table + "s WHERE " + table + "_id IN (" + placeholder + ");"
			vars = InterfaceFromStringArray(objectIDs)
		}
	}
	// execute query
	return ExecuteRowQuery(db, query, vars)
}

func Search(db *sql.DB, table string, name string) (*sql.Rows, error) {

	// prepare select query
	query := "SELECT * FROM " + table + "s WHERE " + table + "_name LIKE CONCAT('%', ?, '%');"
	vars := []interface{}{name}

	// execute query
	return ExecuteRowQuery(db, query, vars)
}

func Resource(db *sql.DB, object string, objectID string, resource string) (*sql.Rows, error) {

	var query string
	// prepare query
	if object == "tag" || resource == "festival" {
		query = "SELECT * FROM " + resource + "s WHERE " + resource + "_id IN (SELECT `associated_" + resource + "` FROM `map_" + resource + "_" + object + "` WHERE `associated_" + object + "`=?);"
	} else {
		query = "SELECT * FROM " + resource + "s WHERE " + resource + "_id IN (SELECT `associated_" + resource + "` FROM `map_" + object + "_" + resource + "` WHERE `associated_" + object + "`=?);"
	}

	vars := []interface{}{objectID}
	// execute query
	return ExecuteRowQuery(db, query, vars)
}

func SetResource(db *sql.DB, object string, objectID string, resource string, resourceID string) error {

	query := "SELECT `map_id` from `map_" + object + "_" + resource + "` WHERE associated_" + object + " =? AND associated_" + resource + "=?;"

	log.Print(query)

	vars := []interface{}{objectID, resourceID}

	log.Print(vars)

	rows, err := ExecuteRowQuery(db, query, vars)
	if err != nil {
		return err
	}

	var mapID string
	log.Print(rows)
	for rows.Next() {
		// scan the link
		err = rows.Scan(&mapID)
		if err != nil {
			return err
		}
	}

	if mapID != "" {
		vars = []interface{}{objectID, resourceID, mapID}
		query = "UPDATE `map_" + object + "_" + resource + "`SET associated_" + object + "=?,associated_" + object + "=?  WHERE map_id=?;"
		_, err := ExecuteQuery(db, query, vars)
		if err != nil {
			return err
		}
		return nil
	} else {
		query = "INSERT INTO `map_" + object + "_" + resource + "` ( `associated_" + object + "` , `associated_" + resource + "` ) VALUES (?,?);"
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

func RemoveResource(db *sql.DB, object string, objectID string, resource string, resourceID string) error {
	// select map id
	query := "SELECT `map_id` from `map_" + object + "_" + resource + "` WHERE associated_" + object + " =? AND associated_" + resource + "=?;"
	vars := []interface{}{objectID, resourceID}
	rows, err := ExecuteRowQuery(db, query, vars)
	if err != nil {
		return err
	}
	if rows != nil {
		var mapID string
		for rows.Next() {
			// scan the link
			err = rows.Scan(&mapID)
			if err != nil {
				return err
			}
		}
		vars = []interface{}{mapID}
		query = "DELETE FROM `map_" + object + "_" + resource + "` WHERE map_id=?;"
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
			return errors.New("remove resource: no rows where affected")
		}
		return nil
	} else {
		return errors.New("resource: there is no resource to remove")
	}
}

func Insert(db *sql.DB, table string, object interface{}) (*sql.Rows, error) {

	fields := DBFields(object)
	placeholder := DBPlaceholder(object)
	vars := DBValues(object)
	query := "INSERT INTO " + table + "s(" + fields + ") VALUES (" + placeholder + ");"
	result, err := ExecuteQuery(db, query, vars)
	if err != nil {
		return nil, err
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	objectID := strconv.FormatInt(insertID, 10)

	return Select(db, table, []string{objectID})
}

func Update(db *sql.DB, table string, objectID string, object interface{}) (*sql.Rows, error) {

	keyValuePairs := DBKeyValuePairs(object)
	vars := DBValues(object)
	vars = append(vars, objectID) // for *table*_id value
	query := "UPDATE " + table + "s SET " + keyValuePairs + " WHERE `" + table + "_id`=?;"

	_, err := ExecuteQuery(db, query, vars)
	if err != nil {
		return nil, err
	}

	return Select(db, table, []string{objectID})
}

func Delete(db *sql.DB, table string, objectID string) error {

	// prepare select query
	query := "DELETE FROM " + table + "s WHERE " + table + "_id=?"
	vars := []interface{}{objectID}
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
		return errors.New("remove resource: no rows where affected")
	}
	return nil
}

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

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return result, nil
}
