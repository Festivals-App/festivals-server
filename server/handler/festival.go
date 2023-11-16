package handler

import (
	"database/sql"
	"net/http"

	servertools "github.com/Festivals-App/festivals-server-tools"
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

func SetEventForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to set event for festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set event for festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetImageForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to set image for festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set image for festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to set link for festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set link for festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetPlaceForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to set place for festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set place for festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func SetTagForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to set tag for festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to set tag for festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image from festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove image from festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove link from festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove link from festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemovePlaceForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove place from festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove place from festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveTagForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove tag from festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to remove tag from festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, []interface{}{})
}

func CreateFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festival, err := Create(db, r, "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to create festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to create festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, festival)
}

func UpdateFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := Update(db, r, "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to update festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to update festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, festivals)
}

func DeleteFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete festival")
		servertools.RespondError(w, http.StatusBadRequest, "failed to delete festival")
		return
	}
	servertools.RespondJSON(w, http.StatusOK, nil)
}
