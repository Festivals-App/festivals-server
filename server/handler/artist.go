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

func GetArtists(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetObjects(db, "artist", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artists")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch artists")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, artists)
}

func GetArtist(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	artists, err := GetObject(db, r, "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, artists)
}

func GetArtistImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "artist", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artist image")
		servertools.RespondError(w, http.StatusInternalServerError, "failed to fetch artist image")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func GetArtistLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetAssociation(db, r, "artist", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artist links")
		servertools.RespondError(w, http.StatusInternalServerError, "failed to fetch artist links")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func GetArtistTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetAssociation(db, r, "artist", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch artist tags")
		servertools.RespondError(w, http.StatusInternalServerError, "failed to fetch artist tags")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, tags)
}

func SetImageForArtist(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to SetImageForArtist.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserArtists, objectID) {
			log.Error().Msg("User is not authorized to SetImageForArtist.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := SetAssociation(db, r, "artist", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to set image for artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set image for artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForArtist(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to SetLinkForArtist.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserArtists, objectID) {
			log.Error().Msg("User is not authorized to SetLinkForArtist.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := SetAssociation(db, r, "artist", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to set link for artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set link for artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetTagForArtist(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to SetTagForArtist.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserArtists, objectID) {
			log.Error().Msg("User is not authorized to SetTagForArtist.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := SetAssociation(db, r, "artist", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to set tag for artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set tag for artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForArtist(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to RemoveImageForArtist.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserArtists, objectID) {
			log.Error().Msg("User is not authorized to RemoveImageForArtist.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := RemoveAssociation(db, r, "artist", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image from artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove image from artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForArtist(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to RemoveLinkForArtist.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserArtists, objectID) {
			log.Error().Msg("User is not authorized to RemoveLinkForArtist.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := RemoveAssociation(db, r, "artist", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove link from artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove link from artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveTagForArtist(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to RemoveTagForArtist.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserArtists, objectID) {
			log.Error().Msg("User is not authorized to RemoveTagForArtist.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := RemoveAssociation(db, r, "artist", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove tag from artist")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove tag from artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func CreateArtist(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.CREATOR && claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to create a tag.")
		servertools.UnauthorizedResponse(w)
		return
	}

	artists, err := Create(db, r, "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to create artist")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if len(artists) != 1 {
		log.Error().Err(err).Msg("failed to retrieve artist after creation")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = registerArtistForUser(claims.UserID, strconv.Itoa(artists[0].(model.Artist).ID), "https://"+claims.Issuer+":22580", config.ServiceKey, validator.Client)
	if err != nil {
		retryToRegisterArtist(artists, validator, claims, config, w)
		return
	}

	servertools.RespondJSON(w, http.StatusOK, artists)
}

func retryToRegisterArtist(artists []interface{}, validator *token.ValidationService, claims *token.UserClaims, config *config.Config, w http.ResponseWriter) {

	time.Sleep(10 * time.Second)

	err := registerArtistForUser(claims.UserID, strconv.Itoa(artists[0].(model.Artist).ID), "https://"+claims.Issuer+":22580", config.ServiceKey, validator.Client)
	if err != nil {
		log.Error().Err(err).Msg("failed to retry to register artist for user")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	servertools.RespondJSON(w, http.StatusOK, artists)
}

func UpdateArtist(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to UpdateArtist.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserArtists, objectID) {
			log.Error().Msg("User is not authorized to UpdateArtist.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	artists, err := Update(db, r, "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to update artist")
		servertools.RespondError(w, http.StatusInternalServerError, "failed to update artist")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, artists)
}

func DeleteArtist(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse object id to DeleteArtist.")
			servertools.RespondError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if !slices.Contains(claims.UserArtists, objectID) {
			log.Error().Msg("User is not authorized to DeleteArtist.")
			servertools.UnauthorizedResponse(w)
			return
		}
	}

	err := Delete(db, r, "artist")
	if err != nil {
		log.Error().Err(err).Msg("failed to delte artist")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
