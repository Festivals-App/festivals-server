package database

import (
	"fmt"
	"reflect"
	"strings"
)

// taken from: https://github.com/jmoiron/sqlx/issues/255
// DBFields reflects on a struct and returns the values of fields with `db` tags and returns the keys.
func DBFields(object interface{}) string {

	v := reflect.ValueOf(object)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var fields []string
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Tag.Get("db")
			if field != "" {
				fields = append(fields, field)
			}
		}
		return "`" + strings.Join(fields, "`,`") + "`"
	}

	panic(fmt.Errorf("DBFields requires a struct, found: %s", v.Kind().String()))
}

// DBPlaceholder reflects on a struct and returns a placeholder string for every field which have a `db` tag.
func DBPlaceholder(object interface{}) string {

	v := reflect.ValueOf(object)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var fields []string
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Tag.Get("db")
			if field != "" {
				fields = append(fields, "?")
			}
		}
		return strings.Join(fields, ",")
	}

	panic(fmt.Errorf("DBFields requires a struct, found: %s", v.Kind().String()))
}

func DBPlaceholderForIDs(ids []int) string {

	placeholderString := ""
	for i := 0; i < len(ids); i++ {
		placeholderString = placeholderString + "?"
		if i != len(ids)-1 {
			placeholderString = placeholderString + ","
		}
	}
	return placeholderString
}

// DBValues reflects on a struct and returns the values of fields which have a `db` tag.
func DBValues(object interface{}) []interface{} {

	v := reflect.ValueOf(object)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var values []interface{}
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Tag.Get("db")
			if field != "" {
				values = append(values, v.Field(i).Interface())
			}
		}
		return values
	}

	panic(fmt.Errorf("DBFields requires a struct, found: %s", v.Kind().String()))
}

// DBValues reflects on a struct and returns the values of fields which have a `db` tag.
func DBKeyValuePairs(object interface{}) string {

	v := reflect.ValueOf(object)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var fields []string
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Tag.Get("db")
			if field != "" {
				fields = append(fields, field)
			}
		}
		return "`" + strings.Join(fields, "`=?,`") + "`=?"
	}

	panic(fmt.Errorf("DBFields requires a struct, found: %s", v.Kind().String()))
}

func InterfaceInt(ints []int) []interface{} {

	b := make([]interface{}, len(ints))
	for i := range ints {
		b[i] = ints[i]
	}
	return b
}
