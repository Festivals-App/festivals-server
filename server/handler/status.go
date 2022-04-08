package handler

import (
	"database/sql"
	"github.com/Festivals-App/festivals-server/server/status"
	"net/http"
)

func GetVersion(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	respondString(w, http.StatusOK, status.VersionString())
}

func GetInfo(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	respondJSON(w, http.StatusOK, status.InfoString())
}

func GetHealth(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	respondCode(w, status.HealthStatus())
}
