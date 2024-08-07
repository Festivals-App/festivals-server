package handler

import (
	"errors"
	"net/http"
	"slices"

	token "github.com/Festivals-App/festivals-identity-server/jwt"
	"github.com/google/uuid"
)

func IsAuthorizedToUseHandler(claims *token.UserClaims, userObjectIDs []int, r *http.Request) error {
	if claims.UserRole != token.ADMIN {
		objectID, err := ObjectID(r)
		if err != nil {
			return err
		}
		if !slices.Contains(userObjectIDs, objectID) {
			return errors.New("user is not authorized to use handler")
		}
	}
	return nil
}

func registerFestivalForUser(userID string, festivalID string, endpoint string, serviceKey string, client *http.Client) error {
	return registerEntityForUser(userID, "festival", festivalID, endpoint, serviceKey, client)
}

func registerArtistForUser(userID string, aritstID string, endpoint string, serviceKey string, client *http.Client) error {
	return registerEntityForUser(userID, "artist", aritstID, endpoint, serviceKey, client)
}

func registerLocationForUser(userID string, locationID string, endpoint string, serviceKey string, client *http.Client) error {
	return registerEntityForUser(userID, "location", locationID, endpoint, serviceKey, client)
}

func registerEventForUser(userID string, eventID string, endpoint string, serviceKey string, client *http.Client) error {
	return registerEntityForUser(userID, "event", eventID, endpoint, serviceKey, client)
}

func registerLinkForUser(userID string, linkID string, endpoint string, serviceKey string, client *http.Client) error {
	return registerEntityForUser(userID, "link", linkID, endpoint, serviceKey, client)
}

func registerImageForUser(userID string, imageID string, endpoint string, serviceKey string, client *http.Client) error {
	return registerEntityForUser(userID, "image", imageID, endpoint, serviceKey, client)
}

func registerPlaceForUser(userID string, placeID string, endpoint string, serviceKey string, client *http.Client) error {
	return registerEntityForUser(userID, "place", placeID, endpoint, serviceKey, client)
}

func registerTagForUser(userID string, tagID string, endpoint string, serviceKey string, client *http.Client) error {
	return registerEntityForUser(userID, "tag", tagID, endpoint, serviceKey, client)
}

func registerEntityForUser(userID string, entity string, entityID string, endpoint string, serviceKey string, client *http.Client) error {

	requestString := endpoint + "/users/" + userID + "/" + entity + "/" + entityID
	request, err := http.NewRequest(http.MethodPost, requestString, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("X-Request-ID", uuid.New().String())
	request.Header.Set("Service-Key", serviceKey)

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to register " + entity + " for user with error: " + http.StatusText(resp.StatusCode))
	}

	return nil
}
