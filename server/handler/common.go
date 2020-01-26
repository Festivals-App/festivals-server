package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {

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

func IDsFromString(ids string) ([]string, error) {

	return strings.Split(ids, ","), nil
}
