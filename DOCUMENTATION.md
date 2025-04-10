<!--suppress ALL -->

<h1 align="center">
    FestivalsAPI Documentation
</h1>

<p align="center">
  <a href="#overview">Overview</a> •
  <a href="#festival-objects">Festivals</a> •
  <a href="#artist-objects">Artists</a> •
  <a href="#location-objects">Locations</a> •
  <a href="#event-objects">Events</a> •
  <a href="#image-objects">Images</a> •
  <a href="#link-objects">Links</a> •
  <a href="#place-objects">Places</a> •
  <a href="#tag-objects">Tags</a>
</p>

### Used Languages

* Documentation: `Markdown`, `HTML`
* Database: `SQL Query Scripts`
* Server Application: `golang`
* Deployment: `bash`

### Authentication

To use the API you need to provide an API key via a custom header or a JWT with your requests authorization header:

```text
Api-Key:<api-key>
Authorization: Bearer <jwt>
```

### Requests

The FestivalsAPI supports the HTTP `GET`, `POST`, `PATCH` and `DELETE` methods.

#### Query Parameter

* `name`  
    The name parameter expects a simple string. `name=^[A-Za-z0-9_.]+$`
* `ids`  
    The ids parameter expects numbers separated by a comma. `ids=1,2,37`
* `include`  
    The include parameter expects the name(s) of the relationship you want the response to include. `include=rel1,rel2,rel3`

### Response

Requests that are handled gracefully by the server will always return a top-level object  
with at least either the`data`or`error`field. The`data`field will always contain an array.  
If the request returns any objects they will be in that array,

```json
{
    "data": [
        {OBJECT},
        {OBJECT},
        {OBJECT}
    ]
}
```

otherwise, an empty array is returned. This is *not* true for the [status](#server-status) API.

```json
{
    "data": []
}
```

If the request specified to include relationships the objects are contained in the`included`field.  
**Included relationships will only work if the request returns only one object.**

```json
{
    "data": [
        {OBJECT}
    ],
    "included": {
        "relationship-1": [
            {OBJECT},
            {OBJECT},    
            {OBJECT}
        ],
        "relationship-2": [
            {OBJECT}
        ]     
    }
}
```

The `error` field will always contain a string with the error message.

```json
{
    "error": "An error occured"
}
```

## Overview

[Status](#server-status)

* GET              `/info`
* GET              `/version`
* POST             `/update`
* GET              `/health`
* GET              `/log`
* GET              `/log/trace`
  
[Festivals](#festival-objects)

* GET, POST             `/festivals` optional `name,ids` query parameter
* GET, PATCH, DELETE    `/festivals/{objectID}`
* GET                   `/festivals/{objectID}/{image|links|place|tags|events}`
* POST, DELETE          `/festivals/{objectID}/{image|links|place|tags|events}/{resourceID}`

[Artists](#artist-objects)

* GET, POST             `/artists`
* GET, PATCH, DELETE    `/artists/{objectID}`
* GET                   `/artists/{objectID}/{image|links|tags}`
* POST, DELETE          `/artists/{objectID}/{image|links|tags}/{resourceID}`

[Locations](#location-objects)

* GET, POST             `/locations`
* GET, PATCH, DELETE    `/locations/{objectID}`
* GET                   `/locations/{objectID}/{image|links|place}`
* POST, DELETE          `/locations/{objectID}/{image|links|place}/{resourceID}`

[Events](#event-objects)

* GET, POST             `/events`
* GET, PATCH, DELETE    `/events/{objectID}`
* GET                   `/events/{objectID}/{image|artist|location}`
* POST, DELETE          `/events/{objectID}/{image|artist|location}/{resourceID}`

[Images](#image-objects)

* GET, POST             `/images`
* GET, PATCH, DELETE    `/images/{objectID}`

[Links](#link-objects)

* GET, POST             `/links`
* GET, PATCH, DELETE    `/links/{objectID}`

[Places](#place-objects)

* GET, POST             `/places`
* GET, PATCH, DELETE    `/places/{objectID}`

[Tags](#tag-objects)

* GET, POST             `/tags`
* GET, PATCH, DELETE    `/tags/{objectID}`

------------------------------------------------------------------------------------

## Server Status

Determine the state of the server.

Info object

```json
{
    "BuildTime":      string,
    "GitRef":         string,
    "Version":        string
}
```

------------------------------------------------------------------------------------

#### GET `/info`

* Authorization: JWT
* Returns
  * Returns the info object
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/version`

* Authorization: JWT
* Returns
  * The version of the server application.
  * Codes `200`/`40x`/`50x`
  * server version as a string `text/plain`
  
------------------------------------------------------------------------------------

#### POST `/update`

Updates to the newest release on github and restarts the service.

* Authorization: JWT
* Returns
  * The version of the server application.
  * Codes `202`/`40x`/`50x`
  * server version as a string `text/plain`

------------------------------------------------------------------------------------

#### GET `/health`

* Authorization: JWT
* Returns
  * Always returns HTTP status code 200
  * Code `200`
  * empty `text/plain`

------------------------------------------------------------------------------------

#### GET `/log`

Returns the service log.

* Authorization: JWT

* Returns
  * Returns a string
  * Codes `200`/`40x`/`50x`
  * empty or `text/plain`

------------------------------------------------------------------------------------

#### GET `/log/trace`

Returns the service trace log.

* Authorization: JWT

* Returns
  * Returns a string
  * Codes `200`/`40x`/`50x`
  * empty or `text/plain`

------------------------------------------------------------------------------------

## Festival Objects

A simple object that represents a festival.

```json
{
    "festival_id":              integer,
    "festival_version":         string,
    "festival_is_valid":        boolean,
    "festival_name":            string,
    "festival_start":           integer,
    "festival_end":             integer,
    "festival_description":     string
}
```

------------------------------------------------------------------------------------

#### GET `/festivals`

Get all festivals.

* Authorization: API Key

* Query Parameter:  
      `name`: Filter result by name  
      `ids` : Filter result by IDs

* Examples:  
      `GET https://localhost:8080/festivals`  
      `GET https://localhost:8080/festivals?name=Stemmwe`  
      `GET https://localhost:8080/festivals?ids=1,8,56`

* Returns
  * Returns the festivals
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/festivals`

Create a new festival

* Authorization: JWT
  
* Examples:  
    `POST https://localhost:8080/festivals`  
    `BODY: {OBJECT}`

* Returns
  * Returns the created festival on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/festivals/{objectID}`

Get the festival with the given `objectID`.

* Authorization: API Key

* Query Parameter:  
    `include`: Include relationships {`image`|`links`|`place`|`tags`|`events`}  
            Note: You need to specify the relationship not the associated object type.

* Examples:  
    `GET https://localhost:8080/festivals/1`  
    `GET https://localhost:8080/festivals/1?include=links,place`

* Returns
  * Returns the festival on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### PATCH `/festivals/{objectID}`

Update the festival with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `PATCH https://localhost:8080/festivals/1`  
    BODY: `{OBJECT}`

* Returns
  * Returns the updated festival on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### DELETE `/festivals/{objectID}`

Delete the festival with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/festivals/1`

* Returns
  * Returns no object.
  * Codes `204`/`40x`/`50x`
  * `data` or `error` field

 ------------------------------------------------------------------------------------

#### GET `/festivals/{objectID}/{image|links|place|tags|events}`

Get the objects that are described by the `{relationship}`.

* Authorization: API Key

* Examples:  
    `GET https://localhost:8080/festivals/1/image`  
             Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns the objects described by the relationship
  * Codes `20x`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/festivals/{objectID}/{image|links|place|tags|events}/{resourceID}`

Adds the object with the given `{resourceID}` to the relationship for the festival with the given `{objectID}`.

* Authorization: JWT
  
* Examples:  
    `POST https://localhost:8080/festivals/1/image/1`
            Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns no object.
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### DELETE `/festivals/{objectID}/{image|links|place|tags|events}/{resourceID}`

Removes the object with the given `{resourceID}` from the relationship for the festival with the given`{objectID}`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/festivals/1/image/1`
            Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns no object.
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

## Artist Objects

A simple object that represents an artist.

```json
{
    "artist_id":            integer,
    "artist_version":       string,
    "artist_name":          string,
    "artist_description":   string
}
```

------------------------------------------------------------------------------------

#### GET `/artists`

Get all artists.

* Authorization: API Key

* Query Parameter:  
      `name`: Filter result by name  
      `ids` : Filter result by IDs

* Examples:  
      `GET https://localhost:8080/artists`  
      `GET https://localhost:8080/artists?name=Beatl`  
      `GET https://localhost:8080/artists?ids=1,8,56`

* Returns
  * Returns the artists
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/artists`

Create a new artist

* Authorization: JWT

* Examples:  
    `POST https://localhost:8080/artists`  
    `BODY: {OBJECT}`

* Returns
  * Returns the created artist on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/artists/{objectID}`

Get the artist with the given `objectID`.

* Authorization: API Key

* Query Parameter:  
    `include`: Include relationships {`image`|`links`|`tags`}  
            Note: You need to specify the relationship not the associated object type.

* Examples:  
    `GET https://localhost:8080/artists/1`  
    `GET https://localhost:8080/artists/1?include=image,tags`

* Returns
  * Returns the artist on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### PATCH `/artists/{objectID}`

Update the artist with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `PATCH https://localhost:8080/artists/1`  
    BODY: `{OBJECT}`

* Returns
  * Returns the updated artist on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### DELETE `/artists/{objectID}`

Delete the artist with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/artists/1`

* Returns
  * Returns no object.
  * Codes `204`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/artists/{objectID}/{image|links|place|tags}`

Get the objects that are described by the `{relationship}`.

* Authorization: API Key

* Examples:  
    `GET https://localhost:8080/artists/1/image`  
             Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns the objects described by the relationship
  * Codes `20x`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/artists/{objectID}/{image|links|tags}/{resourceID}`

Adds the object with the given`{resourceID}`to the relationship for the artist with the given`{objectID}`.

* Authorization: JWT
  
* Examples:  
    `POST https://localhost:8080/artists/1/image/1`
            Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns no object.
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### DELETE `/artists/{objectID}/{image|links|tags}/{resourceID}`

Removes the object with the given`{resourceID}`from the relationship for the artist with the given`{objectID}`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/artists/1/image/1`
            Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns no object.
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

## Location Objects

A simple object that represents a location.

```json
{
    "location_id":              integer,
    "location_version":         string,
    "location_name":            string,
    "location_description":     string,
    "location_accessible":      boolean,
    "location_openair":         boolean
}
```

------------------------------------------------------------------------------------

#### GET `/locations`

Get all locations.

* Authorization: API Key

* Query Parameter:  
      `name`: Filter result by name  
      `ids` : Filter result by IDs

* Examples:  
      `GET https://localhost:8080/locations`  
      `GET https://localhost:8080/locations?name=Beatl`  
      `GET https://localhost:8080/locations?ids=1,8,56`

* Returns
  * Returns the locations
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/locations`

Create a new location

* Authorization: JWT
  
* Examples:  
    `POST https://localhost:8080/locations`  
    `BODY: {OBJECT}`

* Returns
  * Returns the created location on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/locations/{objectID}`

Get the location with the given `objectID`.

* Authorization: API Key

* Query Parameter:  
    `include`: Include relationships {`image`|`links`|`place`}  
            Note: You need to specify the relationship not the associated object type.

* Examples:  
    `GET https://localhost:8080/locations/1`  
    `GET https://localhost:8080/locations/1?include=image,place`

* Returns
  * Returns the location on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### PATCH `/locations/{objectID}`

Update the location with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `PATCH https://localhost:8080/locations/1`  
    BODY: `{OBJECT}`

* Returns
  * Returns the updated location on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### DELETE `/locations/{objectID}`

Delete the location with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/locations/1`

* Returns
  * Returns no object.
  * Codes `204`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/locations/{objectID}/{image|links|place}`

Get the objects that are described by the`{relationship}`.

* Authorization: API Key

* Examples:  
    `GET https://localhost:8080/locations/1/image`  
             Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns the objects described by the relationship
  * Codes `20x`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/locations/{objectID}/{image|links|place}/{resourceID}`

Adds the object with the given`{resourceID}`to the relationship for the location with the given`{objectID}`.

* Authorization: JWT
  
* Examples:  
    `POST https://localhost:8080/locations/1/image/1`
            Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns no object.
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### DELETE `/locations/{objectID}/{image|links|place}/{resourceID}`

Removes the object with the given`{resourceID}`from the relationship for the location with the given`{objectID}`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/locations/1/image/1`
            Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns no object.
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

## Event Objects

A simple object that represents an event.

```json
{
    "event_id":             integer,
    "event_version":        string,
    "event_name":           string,
    "event_description":    string,
    "event_type":           integer,
    "event_start":          integer,
    "event_end":            integer
}
```

------------------------------------------------------------------------------------

#### GET `/events`

Get all events.

* Authorization: API Key
  
* Query Parameter:  
      `ids` : Filter result by IDs

* Examples:  
      `GET https://localhost:8080/events`  
      `GET https://localhost:8080/events?ids=1,8,56`

* Returns
  * Returns the locations
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/events`

Create a new event

* Authorization: JWT
  
* Examples:
    `POST https://localhost:8080/events`  
    `BODY: {OBJECT}`

* Returns
  * Returns the created event on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/events/{objectID}`

Get the event with the given `objectID`.

* Authorization: API Key

* Query Parameter:  
    `include`: Include relationships {`image`|`artist`|`location`}  
            Note: You need to specify the relationship not the associated object type.

* Examples:  
    `GET https://localhost:8080/events/1`  
    `GET https://localhost:8080/events/1?include=image,artist,location`

* Returns
  * Returns the event on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### PATCH `/events/{objectID}`

Update the event with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `PATCH https://localhost:8080/events/1`  
    BODY: `{OBJECT}`

* Returns
  * Returns the updated event on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### DELETE `/events/{objectID}`

Delete the event with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/events/1`

* Returns
  * Returns no object.
  * Codes `204`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/events/{objectID}/{image|artist|location|festival}`

Get the objects that are described by the`{relationship}`.

* Authorization: API Key

* Examples:  
    `GET https://localhost:8080/events/1/artist`  
             Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns the objects described by the relationship
  * Codes `20x`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/events/{objectID}/{image|artist|location}/{resourceID}`

Adds the object with the given`{resourceID}`to the relationship for the event with the given`{objectID}`.

* Authorization: JWT
  
* Examples:  
    `POST https://localhost:8080/events/1/artist/1`
            Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns no object.
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### DELETE `/events/{objectID}/{image|artist|location}/{resourceID}`

Removes the object with the given`{resourceID}`from the relationship for the event with the given`{objectID}`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/events/1/artist/1`
            Note: You need to specify the relationship not the associated object type.

* Returns
  * Returns no object.
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

## Image Objects

A simple object that represents an image.

```json
{
    "image_id":         integer,
    "image_hash":       string,
    "image_comment":    string,
    "image_ref":        string
}
```

------------------------------------------------------------------------------------

#### GET `/images`

Get all images.

* Authorization: API Key

* Query Parameter:  
      `ids` : Filter result by IDs

* Examples:  
      `GET https://localhost:8080/images`  
      `GET https://localhost:8080/images?ids=1,8,56`

* Returns
  * Returns the images
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/images`

Create a new image

* Authorization: JWT
  
* Examples:  
    `POST https://localhost:8080/images`  
    `BODY: {OBJECT}`

* Returns
  * Returns the created image on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/images/{objectID}`

Get the image with the given `objectID`.

* Authorization: API Key

* Examples:  
    `GET https://localhost:8080/images/1`

* Returns
  * Returns the image on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### PATCH `/images/{objectID}`

Update the image with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `PATCH https://localhost:8080/images/1`  
    BODY: `{OBJECT}`

* Returns
  * Returns the updated image on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### DELETE `/images/{objectID}`

Delete the image with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/images/1`

* Returns
  * Returns no object.
  * Codes `204`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

## Link Objects

A simple object that represents a link.

```json
{
    "link_id":      integer,
    "link_version": string,
    "link_url":     string,
    "link_service": integer
}
```

Service type is identified by an integer:

```text
EVTServiceTypeWebsite                       = 0,
EVTServiceTypeEmail                         = 1,
EVTServiceTypePhone                         = 2,
EVTServiceTypeYoutubeProfile                = 8,
EVTServiceTypeYoutubeVideo                  = 3,
EVTServiceTypeYoutubePlaylist               = 9,
EVTServiceTypeYoutubeMusicPlaylist          = 10,
EVTServiceTypeSoundcloudProfile             = 4,
EVTServiceTypeBandcampProfile               = 5,
EVTServiceTypeBandcampTrack                 = 11,
EVTServiceTypeHearthisProfile               = 6,
EVTServiceTypeHearthisEmbededTrack          = 12,
EVTServiceTypeFacebookProfile               = 13,
EVTServiceTypeInstagramProfile              = 14,
EVTServiceTypeSpotifyProfile                = 15,
EVTServiceTypeAppleMusicStoreReferer        = 16,
EVTServiceTypeAppleMusicArtistID            = 17,
EVTServiceTypeShazamProfile                 = 18,
EVTServiceTypeDeezerProfile                 = 19,
EVTServiceTypeTwitterProfile                = 20,
EVTServiceTypeTripadvisorProfile            = 21,
/// An unknown link.
EVTServiceTypeUnknown                       = 7
```

------------------------------------------------------------------------------------

#### GET `/links`

Get all links.

* Authorization: API Key

* Query Parameter:  
      `ids` : Filter result by IDs

* Examples:  
      `GET https://localhost:8080/locations`
      `GET https://localhost:8080/locations?ids=1,8,56`

* Returns
  * Returns the locations
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/links`

Create a new link

* Authorization: JWT
  
* Examples:  
    `POST https://localhost:8080/links`  
    `BODY: {OBJECT}`

* Returns
  * Returns the created image on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/links/{objectID}`

Get the link with the given `objectID`.

* Authorization: API Key

* Examples:  
    `GET https://localhost:8080/links/1`

* Returns
  * Returns the link on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### PATCH `/links/{objectID}`

Update the link with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `PATCH https://localhost:8080/links/1`  
    BODY: `{OBJECT}`

* Returns
  * Returns the updated link on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### DELETE `/links/{objectID}`

Delete the link with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/links/1`

* Returns
  * Returns no object.
  * Codes `204`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

## Place Objects

A simple object that represents a place.

```json
{
    "place_id":                 integer,
    "place_version":            string,
    "place_street":             string,
    "place_zip":                string,
    "place_town":               string,
    "place_street_addition":    string,
    "place_country":            string,
    "place_lat":                float32.
    "place_lon":                float32,
    "place_description":        string
}
```

------------------------------------------------------------------------------------

#### GET `/places`

Get all places.

* Authorization: API Key

* Query Parameter:  
      `ids` : Filter result by IDs

* Examples:  
      `GET https://localhost:8080/places`  
      `GET https://localhost:8080/places?ids=1,8,56`

* Returns
  * Returns the places
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/places`

Create a new place

* Authorization: JWT
  
* Examples:  
    `POST https://localhost:8080/places`  
    `BODY: {OBJECT}`

* Returns
  * Returns the created place on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/places/{objectID}`

Get the place with the given `objectID`.

* Authorization: API Key

* Examples:  
    `GET https://localhost:8080/places/1`

* Returns
  * Returns the place on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### PATCH `/places/{objectID}`

Update the place with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `PATCH https://localhost:8080/places/1`  
    BODY: `{OBJECT}`

* Returns
  * Returns the updated place on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### PATCH `/places/{objectID}`

Delete the place with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/places/1`

* Returns
  * Returns no object.
  * Codes `204`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

## Tag Objects

A simple object that represents a tag.

```json
{
    "tag_id":   integer,
    "tag_name": string
}
```

------------------------------------------------------------------------------------

#### GET `/tags`

Get all tags.

* Authorization: API Key

* Query Parameter:  
      `name`: Filter result by name  
      `ids` : Filter result by IDs

* Examples:  
      `GET https://localhost:8080/tags`  
      `GET https://localhost:8080/tags?name=rock`  
      `GET https://localhost:8080/tags?ids=1,8,56`

* Returns
  * Returns the tags
  * Codes `200`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/tags`

Create a new tag

* Authorization: JWT
  
* Examples:  
    `POST https://localhost:8080/tags`  
    `BODY: {OBJECT}`

* Returns
  * Returns the created tag on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### GET `/tags/{objectID}`

Get the tag with the given `objectID`.

* Authorization: API Key

* Examples:  
    `GET https://localhost:8080/tags/1`

* Returns
  * Returns the tag on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### PATCH `/tags/{objectID}`

Update the tag with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `PATCH https://localhost:8080/tags/1`  
    BODY: `{OBJECT}`

* Returns
  * Returns the updated tag on success.
  * Codes `201`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------

#### DELETE `/tags/{objectID}`

Delete the tag with the given `objectID`.

* Authorization: JWT
  
* Examples:  
    `DELETE https://localhost:8080/tags/1`

* Returns
  * Returns no object.
  * Codes `204`/`40x`/`50x`
  * `data` or `error` field

------------------------------------------------------------------------------------
