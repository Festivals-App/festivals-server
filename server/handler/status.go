package handler

import (
	"database/sql"
	"net/http"

	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/Festivals-App/festivals-server/server/status"
)

func GetVersion(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	servertools.RespondString(w, http.StatusOK, status.VersionString())
}

func GetInfo(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	servertools.RespondJSON(w, http.StatusOK, status.InfoString())
}

func GetHealth(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	servertools.RespondCode(w, status.HealthStatus())
}
