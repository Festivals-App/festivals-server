package handler

import (
	"database/sql"
	"net/http"
	"slices"
	"strconv"
	"time"

	token "github.com/Festivals-App/festivals-identity-server/jwt"
	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/Festivals-App/festivals-server/server/model"
	"github.com/rs/zerolog/log"
)

func GetLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetObjects(db, "link", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch links")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch links")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func GetLink(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetObject(db, r, "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch link")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch link")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func CreateLink(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.CREATOR && claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to create a tag.")
		servertools.UnauthorizedResponse(w)
		return
	}

	links, err := Create(db, r, "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to create link")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create link")
		return
	}

	if len(links) != 1 {
		log.Error().Err(err).Msg("failed to retrieve link after creation")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = registerLinkForUser(claims.UserID, strconv.Itoa(links[0].(model.Link).ID), claims.Issuer, validator.Endpoint, validator.Client)
	if err != nil {
		retryToRegisterLink(links, validator, claims, w)
		return
	}

	servertools.RespondJSON(w, http.StatusOK, links)
}

func retryToRegisterLink(links []interface{}, validator *token.ValidationService, claims *token.UserClaims, w http.ResponseWriter) {

	time.Sleep(10 * time.Second)

	err := registerLinkForUser(claims.UserID, strconv.Itoa(links[0].(model.Link).ID), claims.Issuer, validator.Endpoint, validator.Client)
	if err != nil {
		log.Error().Err(err).Msg("failed to retry to register link for user")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func UpdateLink(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to UpdateLink.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserLinks, objectID) {
			log.Error().Msg("User is not authorized to UpdateLink.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	links, err := Update(db, r, "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to update link")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update link")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func DeleteLink(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to DeleteLink.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserLinks, objectID) {
			log.Error().Msg("User is not authorized to DeleteLink.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := Delete(db, r, "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete link")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete link")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
