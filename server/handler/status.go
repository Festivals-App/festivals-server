package handler

import (
	"database/sql"
	"github.com/rs/zerolog/log"
	"net/http"
	"os/exec"

	"github.com/Festivals-App/festivals-server/server/status"
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

func MakeUpdate(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("/bin/bash", "-c", "sudo /usr/local/festivals-server/run.sh")

	err := cmd.Run()
	if err != nil {
		log.Error().Err(err).Msg("Failed to run update script!")
	}

	respondString(w, http.StatusAccepted, "SHOULD UPDATE")
}
