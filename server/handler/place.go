package handler

import (
	"database/sql"
	"net/http"
	"slices"
	"strconv"
	"time"

	token "github.com/Festivals-App/festivals-identity-server/jwt"
	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/Festivals-App/festivals-server/server/config"
	"github.com/Festivals-App/festivals-server/server/model"
	"github.com/rs/zerolog/log"
)

func GetPlaces(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetObjects(db, "place", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch places")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch places")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, places)
}

func GetPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetObject(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch place")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch place")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, places)
}

func CreatePlace(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.CREATOR && claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to create a tag.")
		servertools.UnauthorizedResponse(w)
		return
	}

	places, err := Create(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to create place")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create place")
		return
	}

	if len(places) != 1 {
		log.Error().Err(err).Msg("failed to retrieve place after creation")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if claims.UserRole != token.ADMIN {
		err = registerPlaceForUser(claims.UserID, strconv.Itoa(places[0].(model.Place).ID), "https://"+claims.Issuer+":22580", config.ServiceKey, validator.Client)
		if err != nil {
			retryToRegisterPlace(places, validator, claims, config, w)
			return
		}
	}

	servertools.RespondJSON(w, http.StatusOK, places)
}

func retryToRegisterPlace(places []interface{}, validator *token.ValidationService, claims *token.UserClaims, config *config.Config, w http.ResponseWriter) {

	time.Sleep(10 * time.Second)

	err := registerPlaceForUser(claims.UserID, strconv.Itoa(places[0].(model.Place).ID), "https://"+claims.Issuer+":22580", config.ServiceKey, validator.Client)
	if err != nil {
		log.Error().Err(err).Msg("failed to retry to register place for user")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, places)
}

func UpdatePlace(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to UpdatePlace.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserPlaces, objectID) {
			log.Error().Msg("User is not authorized to UpdatePlace.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	places, err := Update(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to update place")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update place")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, places)
}

func DeletePlace(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to DeletePlace.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserPlaces, objectID) {
			log.Error().Msg("User is not authorized to DeletePlace.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	if claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to create a tag.")
		servertools.UnauthorizedResponse(w)
		return
	}

	err := Delete(db, r, "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete place")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete place")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
