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

func GetTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetObjects(db, "tag", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch tags")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch tags")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, tags)
}

func GetTag(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetObject(db, r, "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch tag")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch tag")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, tags)
}

func GetTagFestivals(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "tag", "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festivals for tag")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch festivals for tag")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func CreateTag(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.CREATOR && claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to create a tag.")
		servertools.UnauthorizedResponse(w)
		return
	}

	tags, err := Create(db, r, "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to create tag")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create tag")
		return
	}

	if len(tags) != 1 {
		log.Error().Err(err).Msg("failed to retrieve tag after creation")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = registerTagForUser(claims.UserID, strconv.Itoa(tags[0].(model.Tag).ID), "https://"+claims.Issuer+":22580", config.ServiceKey, validator.Client)
	if err != nil {
		retryToRegisterTag(tags, validator, claims, config, w)
		return
	}

	servertools.RespondJSON(w, http.StatusOK, tags)
}

func retryToRegisterTag(tags []interface{}, validator *token.ValidationService, claims *token.UserClaims, config *config.Config, w http.ResponseWriter) {

	time.Sleep(10 * time.Second)

	err := registerTagForUser(claims.UserID, strconv.Itoa(tags[0].(model.Tag).ID), "https://"+claims.Issuer+":22580", config.ServiceKey, validator.Client)
	if err != nil {
		log.Error().Err(err).Msg("failed to retry to register tag for user")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, tags)
}

func UpdateTag(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to UpdateTag.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserTags, objectID) {
			log.Error().Msg("User is not authorized to UpdateTag.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	tags, err := Update(db, r, "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to update tag")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update tag")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, tags)
}

func DeleteTag(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to DeleteTag.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserTags, objectID) {
			log.Error().Msg("User is not authorized to DeleteTag.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := Delete(db, r, "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete tag")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete tag")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
