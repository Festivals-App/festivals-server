package handler

import (
	"database/sql"
	"net/http"

	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/Festivals-App/festivals-server/server/status"
	"github.com/rs/zerolog/log"
)

func MakeUpdate(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	newVersion, err := servertools.RunUpdate(status.ServerVersion, "Festivals-App", "festivals-server", "/usr/local/festivals-server/update.sh")
	if err != nil {
		log.Error().Err(err).Msg("Failed to update")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	servertools.RespondString(w, http.StatusAccepted, newVersion)
}
