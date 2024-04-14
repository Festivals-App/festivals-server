package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"os"

	token "github.com/Festivals-App/festivals-identity-server/jwt"
	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/Festivals-App/festivals-server/server/config"
	"github.com/rs/zerolog/log"
)

func GetLog(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to get server info.")
		servertools.UnauthorizedResponse(w)
		return
	}
	l, err := Log("/var/log/festivals-server/info.log")
	if err != nil {
		log.Error().Err(err).Msg("Failed to get info log.")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	servertools.RespondString(w, http.StatusOK, l)
}

func GetTraceLog(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to get server info.")
		servertools.UnauthorizedResponse(w)
		return
	}
	l, err := Log("/var/log/festivals-server/trace.log.")
	if err != nil {
		log.Error().Err(err).Msg("Failed to get trace log")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	servertools.RespondString(w, http.StatusOK, l)
}

func Log(location string) (string, error) {

	l, err := os.ReadFile(location)
	if err != nil {
		return "", errors.New("Failed to read log file at: '" + location + "' with error: " + err.Error())
	}
	return string(l), nil
}
