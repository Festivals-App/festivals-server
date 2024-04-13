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

func GetLocations(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetObjects(db, "location", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch locations")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch locations")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, locations)
}

func GetLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetObject(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, locations)
}

func GetLocationImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "location", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch image for location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch image for location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func GetLocationLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetAssociation(db, r, "location", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch links for location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch links for location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func GetLocationPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetAssociation(db, r, "location", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch place for location")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch place for location")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, places)
}

func SetImageForLocation(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to SetImageForLocation.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserLocations, objectID) {
			log.Error().Msg("User is not authorized to SetImageForLocation.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := SetAssociation(db, r, "location", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to set image for location")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForLocation(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to SetLinkForLocation.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserLocations, objectID) {
			log.Error().Msg("User is not authorized to SetLinkForLocation.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := SetAssociation(db, r, "location", "link")
	if err != nil {
		log.Error().Err(err).Msg("")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetPlaceForLocation(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to SetPlaceForLocation.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserLocations, objectID) {
			log.Error().Msg("User is not authorized to SetPlaceForLocation.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := SetAssociation(db, r, "location", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to set place for location")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForLocation(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to RemoveImageForLocation.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserLocations, objectID) {
			log.Error().Msg("User is not authorized to RemoveImageForLocation.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := RemoveAssociation(db, r, "location", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image of location")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForLocation(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to RemoveLinkForLocation.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserLocations, objectID) {
			log.Error().Msg("User is not authorized to RemoveLinkForLocation.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := RemoveAssociation(db, r, "location", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image from location")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemovePlaceForLocation(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to RemovePlaceForLocation.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserLocations, objectID) {
			log.Error().Msg("User is not authorized to RemovePlaceForLocation.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := RemoveAssociation(db, r, "location", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove place of location")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func CreateLocation(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.CREATOR && claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to create a location.")
		servertools.UnauthorizedResponse(w)
		return
	}
	locations, err := Create(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to create location")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if len(locations) != 1 {
		log.Error().Err(err).Msg("failed to retrieve location after creation")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = registerLocationForUser(claims.UserID, strconv.Itoa(locations[0].(model.Location).ID), claims.Issuer, validator.Endpoint, validator.Client)
	if err != nil {
		retryToRegisterLocation(locations, validator, claims, w)
		return
	}

	servertools.RespondJSON(w, http.StatusOK, locations)
}

func retryToRegisterLocation(locations []interface{}, validator *token.ValidationService, claims *token.UserClaims, w http.ResponseWriter) {

	time.Sleep(2 * time.Second)

	err := registerLocationForUser(claims.UserID, strconv.Itoa(locations[0].(model.Location).ID), claims.Issuer, validator.Endpoint, validator.Client)
	if err != nil {
		log.Error().Err(err).Msg("failed to retry to register location for user")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, locations)
}

func UpdateLocation(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to UpdateLocation.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserLocations, objectID) {
			log.Error().Msg("User is not authorized to UpdateLocation.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	locations, err := Update(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to update location")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, locations)
}

func DeleteLocation(validator *token.ValidationService, claims *token.UserClaims, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to DeleteLocation.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserLocations, objectID) {
			log.Error().Msg("User is not authorized to DeleteLocation.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := Delete(db, r, "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete location")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
