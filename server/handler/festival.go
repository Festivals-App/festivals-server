package handler

import (
	"database/sql"
	"net/http"

	"github.com/rs/zerolog/log"
)

func GetFestivals(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := GetObjects(db, "festival", nil, r.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festivals")
		respondError(w, http.StatusBadRequest, "failed to fetch festivals")
		return
	}
	respondJSON(w, http.StatusOK, festivals)
}

func GetFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := GetObject(db, r, "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival")
		respondError(w, http.StatusBadRequest, "failed to fetch festival")
		return
	}
	respondJSON(w, http.StatusOK, festivals)
}

func GetFestivalEvents(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	events, err := GetAssociation(db, r, "festival", "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival events")
		respondError(w, http.StatusBadRequest, "failed to fetch festival events")
		return
	}
	respondJSON(w, http.StatusOK, events)
}

func GetFestivalImage(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	images, err := GetAssociation(db, r, "festival", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival image")
		respondError(w, http.StatusBadRequest, "failed to fetch festival image")
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func GetFestivalLinks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	links, err := GetAssociation(db, r, "festival", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival links")
		respondError(w, http.StatusBadRequest, "failed to fetch festival links")
		return
	}
	respondJSON(w, http.StatusOK, links)
}

func GetFestivalPlace(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	places, err := GetAssociation(db, r, "festival", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival place")
		respondError(w, http.StatusBadRequest, "failed to fetch festival place")
		return
	}
	respondJSON(w, http.StatusOK, places)
}

func GetFestivalTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	tags, err := GetAssociation(db, r, "festival", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch festival tags")
		respondError(w, http.StatusBadRequest, "failed to fetch festival tags")
		return
	}
	respondJSON(w, http.StatusOK, tags)
}

func SetEventForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "event")
	if err != nil {
		log.Error().Err(err).Msg("failed to set event for festival")
		respondError(w, http.StatusBadRequest, "failed to set event for festival")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetImageForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to set image for festival")
		respondError(w, http.StatusBadRequest, "failed to set image for festival")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetLinkForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to set link for festival")
		respondError(w, http.StatusBadRequest, "failed to set link for festival")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetPlaceForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to set place for festival")
		respondError(w, http.StatusBadRequest, "failed to set place for festival")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func SetTagForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := SetAssociation(db, r, "festival", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to set tag for festival")
		respondError(w, http.StatusBadRequest, "failed to set tag for festival")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveImageForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "image")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove image from festival")
		respondError(w, http.StatusBadRequest, "failed to remove image from festival")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveLinkForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "link")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove link from festival")
		respondError(w, http.StatusBadRequest, "failed to remove link from festival")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemovePlaceForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "place")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove place from festival")
		respondError(w, http.StatusBadRequest, "failed to remove place from festival")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func RemoveTagForFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := RemoveAssociation(db, r, "festival", "tag")
	if err != nil {
		log.Error().Err(err).Msg("failed to remove tag from festival")
		respondError(w, http.StatusBadRequest, "failed to remove tag from festival")
		return
	}
	respondJSON(w, http.StatusOK, []interface{}{})
}

func CreateFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festival, err := Create(db, r, "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to create festival")
		respondError(w, http.StatusBadRequest, "failed to create festival")
		return
	}
	respondJSON(w, http.StatusOK, festival)
}

func UpdateFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	festivals, err := Update(db, r, "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to update festival")
		respondError(w, http.StatusBadRequest, "failed to update festival")
		return
	}
	respondJSON(w, http.StatusOK, festivals)
}

func DeleteFestival(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	err := Delete(db, r, "festival")
	if err != nil {
		log.Error().Err(err).Msg("failed to delete festival")
		respondError(w, http.StatusBadRequest, "failed to delete festival")
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
