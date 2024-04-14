package handler

import (
	"database/sql"
	"net/http"

	token "github.com/Festivals-App/festivals-identity-server/jwt"
	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/Festivals-App/festivals-server/server/config"
	"github.com/Festivals-App/festivals-server/server/status"
	"github.com/rs/zerolog/log"
)

func GetVersion(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to get server version.")
		servertools.UnauthorizedResponse(w)
		return
	}
	servertools.RespondString(w, http.StatusOK, status.VersionString())
}

func GetInfo(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to get server info.")
		servertools.UnauthorizedResponse(w)
		return
	}
	servertools.RespondJSON(w, http.StatusOK, status.InfoString())
}

func GetHealth(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to get server health.")
		servertools.UnauthorizedResponse(w)
		return
	}
	servertools.RespondCode(w, status.HealthStatus())
}
