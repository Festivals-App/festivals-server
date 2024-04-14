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

func GetEvents(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := GetObjects(db, "event", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch events")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch events")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, events)
}

func GetEvent(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := GetObject(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, events)
}

func GetEventFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := GetAssociation(db, r, "event", "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch festival for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, festivals)
}

func GetEventImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetAssociation(db, r, "event", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch image for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch image for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, artists)
}

func GetEventArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetAssociation(db, r, "event", "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artist for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch artist for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, artists)
}

func GetEventLocation(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	locations, err := GetAssociation(db, r, "event", "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch location for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch location for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, locations)
}

func SetImageForEvent(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to SetImageForEvent.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserEvents, objectID) {
			log.Error().Msg("User is not authorized to SetImageForEvent.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := SetAssociation(db, r, "event", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to set image for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set image for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetArtistForEvent(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to SetArtistForEvent.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserEvents, objectID) {
			log.Error().Msg("User is not authorized to SetArtistForEvent.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := SetAssociation(db, r, "event", "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to set artist for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set artist for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetLocationForEvent(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to SetLocationForEvent.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserEvents, objectID) {
			log.Error().Msg("User is not authorized to SetLocationForEvent.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := SetAssociation(db, r, "event", "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to set location for event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set location for event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForEvent(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to RemoveImageForEvent.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserEvents, objectID) {
			log.Error().Msg("User is not authorized to RemoveImageForEvent.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := RemoveAssociation(db, r, "event", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image from event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove image from event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveArtistForEvent(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to RemoveArtistForEvent.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserEvents, objectID) {
			log.Error().Msg("User is not authorized to RemoveArtistForEvent.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := RemoveAssociation(db, r, "event", "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove artist from event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove artist from event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLocationForEvent(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to RemoveLocationForEvent.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserEvents, objectID) {
			log.Error().Msg("User is not authorized to RemoveLocationForEvent.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := RemoveAssociation(db, r, "event", "location")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove location from event")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove location from event")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func CreateEvent(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.CREATOR && claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to create a tag.")
		servertools.UnauthorizedResponse(w)
		return
	}

	events, err := Create(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to create event")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if len(events) != 1 {
		log.Error().Err(err).Msg("failed to retrieve event after creation")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = registerEventForUser(claims.UserID, strconv.Itoa(events[0].(model.Event).ID), "https://"+claims.Issuer+":22580", config.ServiceKey, validator.Client)
	if err != nil {
		retryToRegisterEvent(events, validator, claims, config, w)
		return
	}

	servertools.RespondJSON(w, http.StatusOK, events)
}

func retryToRegisterEvent(events []interface{}, validator *token.ValidationService, claims *token.UserClaims, config *config.Config, w http.ResponseWriter) {

	time.Sleep(10 * time.Second)

	err := registerEventForUser(claims.UserID, strconv.Itoa(events[0].(model.Event).ID), "https://"+claims.Issuer+":22580", config.ServiceKey, validator.Client)
	if err != nil {
		log.Error().Err(err).Msg("failed to retry to register event for user")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, events)
}

func UpdateEvent(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to UpdateEvent.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserImages, objectID) {
			log.Error().Msg("User is not authorized to UpdateEvent.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	events, err := Update(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to update event")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, events)
}

func DeleteEvent(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to DeleteEvent.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserImages, objectID) {
			log.Error().Msg("User is not authorized to DeleteEvent.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := Delete(db, r, "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete event")
		servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
