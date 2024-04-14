package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	token "github.com/Festivals-App/festivals-identity-server/jwt"
	servertools "github.com/Festivals-App/festivals-server-tools"
	"github.com/Festivals-App/festivals-server/server/config"
	"github.com/Festivals-App/festivals-server/server/model"
	"github.com/rs/zerolog/log"
)

func GetFestivals(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := GetObjects(db, "festival", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festivals")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch festivals")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, festivals)
}

func GetFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := GetObject(db, r, "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, festivals)
}

func GetFestivalEvents(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := GetAssociation(db, r, "festival", "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival events")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch festival events")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, events)
}

func GetFestivalImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "festival", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival image")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch festival image")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, images)
}

func GetFestivalLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetAssociation(db, r, "festival", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival links")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch festival links")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, links)
}

func GetFestivalPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetAssociation(db, r, "festival", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival place")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch festival place")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, places)
}

func GetFestivalTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetAssociation(db, r, "festival", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival tags")
		servertools.RespondError(w, http.StatusBadRequest, "failed to fetch festival tags")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, tags)
}

func SetEventForFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to set event for festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set event for festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetImageForFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to set image for festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set image for festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to set link for festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set link for festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetPlaceForFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to set place for festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set place for festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetTagForFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to set tag for festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set tag for festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image from festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove image from festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove link from festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove link from festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemovePlaceForFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove place from festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove place from festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveTagForFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove tag from festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove tag from festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func CreateFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if claims.UserRole != token.CREATOR && claims.UserRole != token.ADMIN {
		log.Error().Msg("User is not authorized to create a tag.")
		servertools.UnauthorizedResponse(w)
		return
	}

	festivals, err := Create(db, r, "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to create festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create festival")
		return
	}

	if len(festivals) != 1 {
		log.Error().Err(err).Msg("failed to retrieve festival after creation")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = registerFestivalForUser(claims.UserID, strconv.Itoa(festivals[0].(model.Festival).ID), "https://"+claims.Issuer+":22580", config.ServiceKey, validator.Client)
	if err != nil {
		retryToRegisterFestival(festivals, validator, claims, config, w)
		return
	}

	servertools.RespondJSON(w, http.StatusOK, festivals)
}

func retryToRegisterFestival(festivals []interface{}, validator *token.ValidationService, claims *token.UserClaims, config *config.Config, w http.ResponseWriter) {

	time.Sleep(2 * time.Second)

	err := registerFestivalForUser(claims.UserID, strconv.Itoa(festivals[0].(model.Festival).ID), claims.Issuer, config.ServiceKey, validator.Client)
	if err != nil {
		log.Error().Err(err).Msg("failed to retry to register festival for user")
		servertools.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	servertools.RespondJSON(w, http.StatusOK, festivals)
}

func UpdateFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := Update(db, r, "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to update festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, festivals)
}

func DeleteFestival(validator *token.ValidationService, claims *token.UserClaims, config *config.Config, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
