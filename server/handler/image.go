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

func GetImages(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetObjects(db, "image", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch images")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch images")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func GetImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetObject(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch image")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch image")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func CreateImage(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.CREATOR && claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to create a tag.")
		servertools.UnauthorizedResponse(w)
		return
	}

	images, err := Create(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to create image")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create image")
		return
	}

	if len(images) != 1 {
		log.Error().Err(err).Msg("failed to retrieve image after creation")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = registerImageForUser(claims.UserID, strconv.Itoa(images[0].(model.Image).ID), claims.Issuer, validator.Endpoint, validator.Client)
	if err != nil {
		retryToRegisterImage(images, validator, claims, w)
		return
	}

	servertools.RespondJSON(w, http.StatusOK, images)
}

func retryToRegisterImage(images []interface{}, validator *token.ValidationService, claims *token.UserClaims, w http.ResponseWriter) {

	time.Sleep(10 * time.Second)

	err := registerImageForUser(claims.UserID, strconv.Itoa(images[0].(model.Image).ID), claims.Issuer, validator.Endpoint, validator.Client)
	if err != nil {
		log.Error().Err(err).Msg("failed to retry to register image for user")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func UpdateImage(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to UpdateImage.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserImages, objectID) {
			log.Error().Msg("User is not authorized to UpdateImage.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	images, err := Update(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to update image")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update image")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func DeleteImage(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to DeleteImage.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserImages, objectID) {
			log.Error().Msg("User is not authorized to DeleteImage.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := Delete(db, r, "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete image")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete image")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
