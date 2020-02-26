package database

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Select(db *sql.DB, table string, objectIDs []int) (*sql.Rows, error) {

	var query string
	var vars []int
	if len(objectIDs) == 0 {
		query = "SELECT * FROM " + table + "s;"
		vars = []int{}
	} else {
		placeholder := DBPlaceholderForIDs(objectIDs)
		query = "SELECT * FROM " + table + "s WHERE " + table + "_id IN (" + placeholder + ");"
		vars = objectIDs
	}
	return ExecuteRowQuery(db, query, InterfaceInt(vars))
}

func Search(db *sql.DB, table string, name string) (*sql.Rows, error) {

	query := "SELECT * FROM " + table + "s WHERE " + table + "_name LIKE CONCAT('%', ?, '%');"
	vars := []interface{}{name}
	return ExecuteRowQuery(db, query, vars)
}

func Resource(db *sql.DB, object string, objectID int, resource string) (*sql.Rows, error) {

	var query string
	if object == "tag" || resource == "festival" {
		query = "SELECT * FROM " + resource + "s WHERE " + resource + "_id IN (SELECT `associated_" + resource + "` FROM `map_" + resource + "_" + object + "` WHERE `associated_" + object + "`=?);"
	} else {
		query = "SELECT * FROM " + resource + "s WHERE " + resource + "_id IN (SELECT `associated_" + resource + "` FROM `map_" + object + "_" + resource + "` WHERE `associated_" + object + "`=?);"
	}
	vars := []interface{}{objectID}
	return ExecuteRowQuery(db, query, vars)
}

func SetResource(db *sql.DB, object string, objectID int, resource string, resourceID int) error {

	query := "SELECT `map_id` from `map_" + object + "_" + resource + "` WHERE associated_" + object + " =?;"

	log.Print(query)

	vars := []interface{}{objectID}
	rows, err := ExecuteRowQuery(db, query, vars)
	if err != nil {
		return err
	}
	var mapID string
	for rows.Next() {
		err = rows.Scan(&mapID)
		if err != nil {
			return err
		}
	}

	log.Print("'" + mapID + "'")

	if mapID != "" {
		log.Print("update association")
		vars = []interface{}{objectID, resourceID, mapID}
		query = "UPDATE `map_" + object + "_" + resource + "` SET associated_" + resource + "=? WHERE map_id=?;"
		log.Print(query)
		vars := []interface{}{resourceID, mapID}
		_, err := ExecuteQuery(db, query, vars)
		if err != nil {
			return err
		}
		return nil
	} else {
		log.Print("create association")
		query = "INSERT INTO `map_" + object + "_" + resource + "` ( `associated_" + object + "` , `associated_" + resource + "` ) VALUES (?,?);"
		vars := []interface{}{objectID, resourceID}
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

func RemoveResource(db *sql.DB, object string, objectID int, resource string, resourceID int) error {

	query := "SELECT `map_id` FROM `map_" + object + "_" + resource + "` WHERE associated_" + object + " =? AND associated_" + resource + "=?;"
	vars := []interface{}{objectID, resourceID}
	rows, err := ExecuteRowQuery(db, query, vars)
	if err != nil {
		return err
	}
	if rows != nil {
		var mapID string
		for rows.Next() {
			err = rows.Scan(&mapID)
			if err != nil {
				return err
			}
		}
		vars = []interface{}{mapID}
		query = "DELETE FROM `map_" + object + "_" + resource + "` WHERE map_id=?;"
		result, err := ExecuteQuery(db, query, vars)
		if err != nil {
			return err
		}
		numOfAffectedRows, err := result.RowsAffected()
		if err != nil {
			return err
		}
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
	return Select(db, table, []int{int(insertID)})
}

func Update(db *sql.DB, table string, objectID int, object interface{}) (*sql.Rows, error) {

	keyValuePairs := DBKeyValuePairs(object)
	vars := DBValues(object)
	vars = append(vars, objectID) // for *table*_id value
	query := "UPDATE " + table + "s SET " + keyValuePairs + " WHERE `" + table + "_id`=?;"
	_, err := ExecuteQuery(db, query, vars)
	if err != nil {
		return nil, err
	}
	return Select(db, table, []int{objectID})
}

func Delete(db *sql.DB, table string, objectID int) error {

	query := "DELETE FROM " + table + "s WHERE " + table + "_id=?"
	vars := []interface{}{objectID}
	result, err := ExecuteQuery(db, query, vars)
	if err != nil {
		return err
	}
	numOfAffectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if numOfAffectedRows != 1 {
		return errors.New("remove resource: no rows where affected")
	}
	return nil
}

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
