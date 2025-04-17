<h1 align="center">
    FestivalsAPI Documentation
</h1>

<p align="center">
  <a href="#overview">Overview</a> •
  <a href="#server-status">Server Status</a> •
  <a href="#festival-route">Festivals</a> •
  <a href="#artist-route">Artists</a> •
  <a href="#location-route">Locations</a> •
  <a href="#event-route">Events</a> •
  <a href="#image-route">Images</a> •
  <a href="#link-route">Links</a> •
  <a href="#place-route">Places</a> •
  <a href="#tag-route">Tags</a>
</p>

### Used Languages

* Documentation: `Markdown`, `HTML`
* Database: `SQL Query Scripts`
* Server Application: `golang`
* Deployment: `bash`

### Authentication

To authenticate to the `FestivalsAPI` you need to either provide an API key via a custom header or a JWT
with your requests authorization header:

```text
Api-Key: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
Authorization: Bearer <JWT>
```

If you have the authorization to call the given endpoint is determined by your
[user role](https://github.com/Festivals-App/festivals-identity-server/blob/master/auth/user.go).

#### Making a request with curl

```bash
curl -H "X-Request-ID: <uuid>" -H "Authorization: Bearer <JWT>" --cacert ca.crt --cert client.crt --key client.key https://api.festivalsapp.home/info
```

### Query Parameter

* `name`  
    The name parameter expects a simple string. `name=^[A-Za-z0-9_.]+$`
* `ids`  
    The ids parameter expects numbers separated by a comma. `ids=1,2,37`
* `include`  
    The include parameter expects the name(s) of the relationship you want the response to include. `include=rel1,rel2,rel3`

### Response

Requests that are handled gracefully by the server will always return a top-level object  
with at least either the `data` or `error` field. The `data` field will always contain an array.  
Except for the info endopoint, if the request returns any objects they will be in that array,

```json
{
    "data": [
        {"OBJECT"},
        {"OBJECT"},
        {"OBJECT"}
    ]
}
```

otherwise, an empty array is returned.

```json
{
    "data": []
}
```

If the request specified to include relationships the objects are contained in the `included` field.  
**Included relationships will only work if the request returns only one object.**

```json
{
    "data": [
        {"OBJECT"}
    ],
    "included": {
        "relationship-1": [
            {"OBJECT"},
            {"OBJECT"},    
            {"OBJECT"}
        ],
        "relationship-2": [
            {"OBJECT"}
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

[Server Status](#server-status)

* GET              `/info`
* GET              `/version`
* POST             `/update`
* GET              `/health`
* GET              `/log`
* GET              `/log/trace`
  
[Festivals](#festival-route)

* GET, POST             `/festivals` optional `name,ids` query parameter
* GET, PATCH, DELETE    `/festivals/{objectID}`
* GET                   `/festivals/{objectID}/{image|links|place|tags|events}`
* POST, DELETE          `/festivals/{objectID}/{image|links|place|tags|events}/{resourceID}`

[Artists](#artist-route)

* GET, POST             `/artists`
* GET, PATCH, DELETE    `/artists/{objectID}`
* GET                   `/artists/{objectID}/{image|links|tags}`
* POST, DELETE          `/artists/{objectID}/{image|links|tags}/{resourceID}`

[Locations](#location-route)

* GET, POST             `/locations`
* GET, PATCH, DELETE    `/locations/{objectID}`
* GET                   `/locations/{objectID}/{image|links|place}`
* POST, DELETE          `/locations/{objectID}/{image|links|place}/{resourceID}`

[Events](#event-route)

* GET, POST             `/events`
* GET, PATCH, DELETE    `/events/{objectID}`
* GET                   `/events/{objectID}/{image|artist|location}`
* POST, DELETE          `/events/{objectID}/{image|artist|location}/{resourceID}`

[Images](#image-route)

* GET, POST             `/images`
* GET, PATCH, DELETE    `/images/{objectID}`

[Links](#link-route)

* GET, POST             `/links`
* GET, PATCH, DELETE    `/links/{objectID}`

[Places](#place-route)

* GET, POST             `/places`
* GET, PATCH, DELETE    `/places/{objectID}`

[Tags](#tag-route)

* GET, POST             `/tags`
* GET, PATCH, DELETE    `/tags/{objectID}`

------------------------------------------------------------------------------------

## Server Status

The **server status routes** serve status-related information and are available at:

It is commonly used for health checks, CI/CD diagnostics, or runtime introspection. This route uses
a `server-info` object containing metadata about the currently running binary, such as build time,
Git reference, service name, and version.

**`server-info`** object

```json
{
  "BuildTime": "string",
  "GitRef": "string",
  "Service": "string",
  "Version": "string"
}
```

| Field      | Description                                                                 |
|------------|-----------------------------------------------------------------------------|
| `BuildTime` | Timestamp of the binary build. Format: `Sun Apr 13 13:55:44 UTC 2025`       |
| `GitRef`    | Git reference used for the build. Format: `refs/tags/v2.2.0` [🔗 Git Docs](https://git-scm.com/book/en/v2/Git-Internals-Git-References) |
| `Service`   | Service identifier. Matches a defined [Service type](https://github.com/Festivals-App/festivals-server-tools/blob/main/heartbeattools.go) |
| `Version`   | Version tag of the deployed binary. Format: `v2.2.0`                        |

> In production builds, these values are injected at build time and reflect the deployment source and context.

------------------------------------------------------------------------------------

#### GET `/info`

Returns the `server-info`.

Example:  
  `GET https://api.festivalsapp.home/info`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN`.

**Response**

* `data` or `error` field
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/version`

Returns the release version of the server.

> In production builds this will have the format `v2.2.0` but
for manual builds this will may be `development`.

Example:  
  `GET https://api.festivalsapp.home/version`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN`.

**Response**

* Server version as a string `text/plain`.
* Codes `200`/`40x`/`50x`
  
------------------------------------------------------------------------------------

#### POST `/update`

Updates to the newest release on github and restarts the service.

Example:  
  `POST https://api.festivalsapp.home/update`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN`.

**Response**

* The current version and the version the server is updated to as a string `text/plain`. Format: `v2.1.3 => v2.2.0`
* Codes `202`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/health`

A simple health check endpoint that returns a `200 OK` status if the service is running and able to respond.

Example:  
  `GET https://api.festivalsapp.home/health`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN`.

**Response**

* Always returns `200 OK`

------------------------------------------------------------------------------------

#### GET `/log`

Returns the info log file as a string, containing all log messages except trace log entries.
See [loggertools](https://github.com/Festivals-App/festivals-server-tools/blob/main/DOCUMENTATION.md#loggertools) for log format.

Example:  
  `GET https://api.festivalsapp.home/log`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN`.

**Response**

* Returns a string as `text/plain`
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/log/trace`

Returns the trace log file as a string, containing all remote calls to the server.
See [loggertools](https://github.com/Festivals-App/festivals-server-tools/blob/main/DOCUMENTATION.md#loggertools) for log format.

Example:  
  `GET https://api.festivalsapp.home/log/trace`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN`.

**Response**

* Returns a string as `text/plain`
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

## Festival Route

The **festival route** provides endpoints to retrieve and edit festivals and associated entities.
This route uses a `festival` object describing a festival in the FestivalsApp.

**`festival`** object

```json
{
  "festival_id":          "int",
  "festival_version":     "string",
  "festival_is_valid":    "string",
  "festival_name":        "string",
  "festival_start":       "string",
  "festival_end":         "string",
  "festival_description": "string",
  "festival_price":       "string"
}
```

| Field                  | Description                                                          |
|------------------------|----------------------------------------------------------------------|
| `festival_id`          | The id of the festival.                                              |
| `festival_version`     | The version of the festival. Format: 2024-07-02T22:41:02Z            |
| `festival_is_valid`    | Boolean value indicating if the festival should be visible to users. |
| `festival_name`        | The festival name.                                                   |
| `festival_start`       | The start date of the festival. Format: Seconds till UNIX Time.      |  
| `festival_end`         | The end date of the festival. Format: Seconds till UNIX Time.        |
| `festival_description` | The description of the festival.                                     |
| `festival_price`       | The price description of the festival.                               |

------------------------------------------------------------------------------------

#### GET `/festivals`

Gets all festivals as a list of `festival`s.

Query Parameter:
  `name`: Filter result by name
  `ids` : Filter result by IDs

Examples:  
  `GET https://api.festivalsapp.home/festivals`  
  `GET https://api.festivalsapp.home/festivals?name=Stemmwe`  
  `GET https://api.festivalsapp.home/festivals?ids=1,8,56`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### POST `/festivals`

Creates a new festival with the provided `festival` and returns it as a `festival`.

Examples:  
  `POST https://api.festivalsapp.home/festivals`  
  `BODY: {festival}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/festivals/{objectID}`

Gets the festival with the given `objectID` as a `festival`.

Examples:  
  `GET https://api.festivalsapp.home/festivals/1`  
  `GET https://api.festivalsapp.home/festivals/1?include=links,place`

**Authorization**
Requires a valid `Api-Key`.

**Query Parameter**  

  `include`: Include relationships {`image`|`links`|`place`|`tags`|`events`}  
    Note: You need to specify the relationship not the associated object type.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### PATCH `/festivals/{objectID}`

Updates the festival with the given `objectID` and return it as a `festival`.
  
Examples:  
  `PATCH https://api.festivalsapp.home/festivals/1`  
  `BODY: {festival}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/festivals/{objectID}`

Deletes the festival with the given `objectID`.

Examples:  
  `DELETE https://api.festivalsapp.home/festivals/1`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `204 No Content` on success and `error` on failure.
* Codes `204`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/festivals/{objectID}/{image|links|place|tags|events}`

Gets the objects that are described by the `{relationship}` and returns it as an `{image|link|place|tag|event}`.

Examples:  
  `GET https://localhost:8080/festivals/1/image`  
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `20x`/`40x`/`50x`

------------------------------------------------------------------------------------

#### POST `/festivals/{objectID}/{image|links|place|tags|events}/{resourceID}`

Adds the object with the given `{resourceID}` to the relationship for the festival with the given `{objectID}`.

Examples:  
  `POST https://localhost:8080/festivals/1/image/1`
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `200 OK` on success and `error` on failure.
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/festivals/{objectID}/{image|links|place|tags|events}/{resourceID}`

Removes the object with the given `{resourceID}` from the relationship for the festival with the given`{objectID}`.

Examples:  
  `DELETE https://localhost:8080/festivals/1/image/1`
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `204 No Content` on success and `error` on failure.
* Codes `204`/`40x`/`50x`

------------------------------------------------------------------------------------

## Artist Route

The **artist route** provides endpoints to retrieve and edit artists and associated entities.
This route uses an `artist` object describing an artist in the FestivalsApp.

**`artist`** object

```json
{
    "artist_id":            "int",
    "artist_version":       "string",
    "artist_name":          "string",
    "artist_description":   "string"
}
```

| Field                  | Description                                                |
|------------------------|------------------------------------------------------------|
| `artist_id`            | The id of the artist.                                      |
| `artist_version`       | The version of the artist. Format: 2024-07-02T22:41:02Z    |
| `artist_name`          | The artist name.                                           |
| `artist_description`   | The description of the artist.                             |

------------------------------------------------------------------------------------

#### GET `/artists`

Gets all artists as a list of `atist`s.

Query Parameter:  
  `name`: Filter result by name  
  `ids` : Filter result by IDs

Examples:  
  `GET https://localhost:8080/artists`  
  `GET https://localhost:8080/artists?name=Beatl`  
  `GET https://localhost:8080/artists?ids=1,8,56`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### POST `/artists`

Creates a new artist with the provided `artist` and returns it as an `artist`.

Examples:  
  `POST https://localhost:8080/artists`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/artists/{objectID}`

Gets the artist with the given `objectID` and returns it as an `artist`.

Query Parameter:  
  `include`: Include relationships {`image`|`links`|`tags`}  
    Note: You need to specify the relationship not the associated object type.

Examples:  
  `GET https://localhost:8080/artists/1`  
  `GET https://localhost:8080/artists/1?include=image,tags`

**Authorization**
Requires a valid `Api-Key`.

**Response**
  
* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### PATCH `/artists/{objectID}`

Updates the artist with the given `objectID` and returns it as an `artist`.

Examples:  
  `PATCH https://localhost:8080/artists/1`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/artists/{objectID}`

Deletes the artist with the given `objectID`.

Examples:  
  `DELETE https://localhost:8080/artists/1`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `204 No Content` on success and `error` on failure.
* Codes `204`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/artists/{objectID}/{image|links|place|tags}`

Gets the objects that are described by the `{relationship}` and returns it as an `{image|link|place|tag}`.

Examples:  
  `GET https://localhost:8080/artists/1/image`  
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `20x`/`40x`/`50x`

------------------------------------------------------------------------------------

#### POST `/artists/{objectID}/{image|links|tags}/{resourceID}`

Adds the object with the given`{resourceID}`to the relationship for the artist with the given`{objectID}`.

Examples:  
  `POST https://localhost:8080/artists/1/image/1`
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `200 OK` on success and `error` on failure.
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/artists/{objectID}/{image|links|tags}/{resourceID}`

Removes the object with the given`{resourceID}`from the relationship for the artist with the given`{objectID}`.

Examples:  
  `DELETE https://localhost:8080/artists/1/image/1`
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `204 No Content` on success and `error` on failure.
* Codes `204`/`40x`/`50x`

------------------------------------------------------------------------------------

## Location Route

The **location route** provides endpoints to retrieve and edit locations and associated entities.
This route uses a `location` object describing a location in the FestivalsApp.

**`location`** object

```json
{
    "location_id":              "int",
    "location_version":         "string",
    "location_name":            "string",
    "location_description":     "string",
    "location_accessible":      "boolean",
    "location_openair":         "boolean"
}
```

| Field                  | Description                                                  |
|------------------------|--------------------------------------------------------------|
| `location_id`          | The id of the location.                                      |
| `location_version`     | The version of the location. Format: 2024-07-02T22:41:02Z    |
| `location_name`        | The location name.                                           |
| `location_description` | The description of the location.                             |
| `location_accessible`  | Boolean value indicating if the location is accessible.      |
| `location_openair`     | Boolean value indicating if the festival is open-air.        |

------------------------------------------------------------------------------------

#### GET `/locations`

Gets all locations as a list of `locations`s.

Query Parameter:  
  `name`: Filter result by name  
  `ids` : Filter result by IDs

Examples:  
  `GET https://localhost:8080/locations`  
  `GET https://localhost:8080/locations?name=Beatl`  
  `GET https://localhost:8080/locations?ids=1,8,56`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### POST `/locations`

Creates a new location with the provided `location` and returns it as a `location`.
  
Examples:  
  `POST https://localhost:8080/locations`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/locations/{objectID}`

Gets the location with the given `objectID` and returns it as an `location`.

Query Parameter:  
  `include`: Include relationships {`image`|`links`|`place`}  
    Note: You need to specify the relationship not the associated object type.

Examples:  
  `GET https://localhost:8080/locations/1`  
  `GET https://localhost:8080/locations/1?include=image,place`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### PATCH `/locations/{objectID}`

Updates the location with the given `objectID`and returns it as an `location`.

Examples:  
  `PATCH https://localhost:8080/locations/1`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/locations/{objectID}`

Deletes the location with the given `objectID`.

Examples:  
  `DELETE https://localhost:8080/locations/1`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**
  
* Returns `204 No Content` on success and `error` on failure.
* Codes `204`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/locations/{objectID}/{image|links|place}`

Gets the objects that are described by the `{relationship}` and returns it as an `{image|links|place}`.

Examples:  
  `GET https://localhost:8080/locations/1/image`  
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `20x`/`40x`/`50x`

------------------------------------------------------------------------------------

#### POST `/locations/{objectID}/{image|links|place}/{resourceID}`

Adds the object with the given`{resourceID}`to the relationship for the location with the given`{objectID}`.

Examples:  
  `POST https://localhost:8080/locations/1/image/1`
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `200 OK` on success and `error` on failure.
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/locations/{objectID}/{image|links|place}/{resourceID}`

Removes the object with the given`{resourceID}`from the relationship for the location with the given`{objectID}`.

Examples:  
  `DELETE https://localhost:8080/locations/1/image/1`
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `204 No Content` on success and `error` on failure.
* Codes `204`/`40x`/`50x`

------------------------------------------------------------------------------------

## Event Route

The **event route** provides endpoints to retrieve and edit events and associated entities.
This route uses an `event` object describing an event in the FestivalsApp.

**`event`** object

```json
{
    "event_id":             "int",
    "event_version":        "string",
    "event_name":           "string",
    "event_description":    "string",
    "event_type":           "int",
    "event_start":          "int",
    "event_end":            "int",
    "event_is_scheduled":   "bool",
    "event_has_timeslot":   "bool"
}
```

| Field                | Description                                                     |
|----------------------|-----------------------------------------------------------------|
| `event_id`           | The id of the event.                                            |
| `event_version`      | The version of the event. Format: 2024-07-02T22:41:02Z          |
| `event_name`         | The event name.                                                 |
| `event_description`  | The description of the event.                                   |
| `event_type`         | The type of the event. Defaults to type music.                  |
| `event_start`        | The start date of the event. Format: Seconds till UNIX Time.    |
| `event_end`          | The end date of the event. Format: Seconds till UNIX Time.      |
| `event_is_scheduled` | Boolean value indicating if the event has a scheduled date.     |
| `event_has_timeslot` | Boolean value indicating if the event has a scheduled timeslot. |

------------------------------------------------------------------------------------

#### GET `/events`

Gets all events as a list of `events`s.
  
Query Parameter:  
  `ids` : Filter result by IDs

Examples:  
  `GET https://localhost:8080/events`  
  `GET https://localhost:8080/events?ids=1,8,56`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### POST `/events`

Creates a new event with the provided `event` and returns it as an `event`.

Examples:
  `POST https://localhost:8080/events`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/events/{objectID}`

Gets the event with the given `objectID` and returns it as an `event`.

Query Parameter:  
  `include`: Include relationships {`image`|`artist`|`location`}  
    Note: You need to specify the relationship not the associated object type.

Examples:  
  `GET https://localhost:8080/events/1`  
  `GET https://localhost:8080/events/1?include=image,artist,location`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### PATCH `/events/{objectID}`

Updates the event with the given `objectID` and returns it as an `event`.

Examples:  
  `PATCH https://localhost:8080/events/1`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/events/{objectID}`

Delete the event with the given `objectID`.

Examples:  
  `DELETE https://localhost:8080/events/1`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `204 No Content` on success and `error` on failure.
* Codes `204`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/events/{objectID}/{image|artist|location|festival}`

Gets the objects that are described by the `{relationship}` and return it as an `{image|artist|location|festival}`.

Examples:  
  `GET https://localhost:8080/events/1/artist`  
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `Api-Key`.

**Response**
  
* `data` or `error` field
* Codes `20x`/`40x`/`50x`

------------------------------------------------------------------------------------

#### POST `/events/{objectID}/{image|artist|location}/{resourceID}`

Adds the object with the given`{resourceID}`to the relationship for the event with the given`{objectID}`.

Examples:  
  `POST https://localhost:8080/events/1/artist/1`
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `200 OK` on success and `error` on failure.
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/events/{objectID}/{image|artist|location}/{resourceID}`

Removes the object with the given`{resourceID}`from the relationship for the event with the given`{objectID}`.

Examples:  
  `DELETE https://localhost:8080/events/1/artist/1`
    Note: You need to specify the relationship not the associated object type.

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `204 No Content` on success and `error` on failure.
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

## Image Route

The **image route** provides endpoints to retrieve and edit images and associated entities.
This route uses an `image` object describing an image in the FestivalsApp.

**`image`** object

```json
{
    "image_id":         "int",
    "image_hash":       "string",
    "image_comment":    "string",
    "image_ref":        "string"
}
```

| Field           | Description                     |
|-----------------|---------------------------------|
| `image_id`      | The id of the image.            |
| `image_hash`    | The hash of the image.          |
| `image_comment` | The comment of the image.       |
| `image_ref`     | The referer link of the image.  |

------------------------------------------------------------------------------------

#### GET `/images`

Gets all images as a list of `images`s.

Query Parameter:  
  `ids` : Filter result by IDs

Examples:  
  `GET https://localhost:8080/images`  
  `GET https://localhost:8080/images?ids=1,8,56`

**Authorization**
Requires a valid `Api-Key`.

**Response**
  
* Codes `200`/`40x`/`50x`
* `data` or `error` field

------------------------------------------------------------------------------------

#### POST `/images`

Creats a new image with the provided `image` and returns it as an `image`.

Examples:  
  `POST https://localhost:8080/images`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/images/{objectID}`

Gets the image with the given `objectID` and return it as an `image`.

Examples:  
  `GET https://localhost:8080/images/1`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### PATCH `/images/{objectID}`

Updates the image with the given `objectID` and return it as an `image`.

Examples:  
  `PATCH https://localhost:8080/images/1`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/images/{objectID}`

Deletes the image with the given `objectID`.

Examples:  
  `DELETE https://localhost:8080/images/1`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `204 No Content` on success and `error` on failure.
* Codes `204`/`40x`/`50x`

------------------------------------------------------------------------------------

## Link Route

The **link route** provides endpoints to retrieve and edit links and associated entities.
This route uses a `link` object describing a link in the FestivalsApp.

**`link`** object

```json
{
    "link_id":      "int",
    "link_version": "string",
    "link_url":     "string",
    "link_service": "int"
}
```

| Field          | Description                   |
|----------------|-------------------------------|
| `link_id`      | The id of the link.           |
| `link_version` | The version of the link.      |
| `link_url`     | The url of the link.          |
| `link_service` | The service type of the link. |

The **service type** is identified by an int:

```golang
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

Gets all links as a list of `links`s.

Query Parameter:  
  `ids` : Filter result by IDs

Examples:  
  `GET https://localhost:8080/locations`
  `GET https://localhost:8080/locations?ids=1,8,56`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### POST `/links`

Creates a new link with the provided `link` and returns it as a `link`.

Examples:  
  `POST https://localhost:8080/links`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/links/{objectID}`

Gets the link with the given `objectID` and return it as a `link`.

Examples:  
  `GET https://localhost:8080/links/1`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### PATCH `/links/{objectID}`

Updates the link with the given `objectID`  and return it as a `link`.

Examples:  
  `PATCH https://localhost:8080/links/1`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/links/{objectID}`

Deletes the link with the given `objectID`.

Examples:  
  `DELETE https://localhost:8080/links/1`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `204 No Content` on success and `error` on failure.
* Codes `204`/`40x`/`50x`

------------------------------------------------------------------------------------

## Place Route

The **place route** provides endpoints to retrieve and edit places and associated entities.
This route uses a `place` object describing a place in the FestivalsApp.

**`place`** object

```json
{
    "place_id":                 "int",
    "place_version":            "string",
    "place_street":             "string",
    "place_zip":                "string",
    "place_town":               "string",
    "place_street_addition":    "string",
    "place_country":            "string",
    "place_lat":                "float32",
    "place_lon":                "float32",
    "place_description":        "string"
}
```

| Field                   | Description                       |
|-------------------------|-----------------------------------|
| `place_id`              | The id of the place.              |
| `place_version`         | The version of the place.         |
| `place_street`          | The street of the place.          |
| `place_zip`             | The zip code of the place.        |
| `place_town`            | The town of the place.            |
| `place_street_addition` | The street addition of the place. |
| `place_country`         | The country of the place.         |
| `place_lat`             | The latitude of the place.        |
| `place_lon`             | The longitude of the place.       |
| `place_description`     | The description of the place.     |

------------------------------------------------------------------------------------

#### GET `/places`

Gets all places as a list of `places`s.

Query Parameter:  
  `ids` : Filter result by IDs

Examples:  
  `GET https://localhost:8080/places`  
  `GET https://localhost:8080/places?ids=1,8,56`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### POST `/places`

Creates a new place with the provided `place` and returns it as a `place`.

Examples:  
  `POST https://localhost:8080/places`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/places/{objectID}`

Gets the place with the given `objectID` and return it as a `place`.

Examples:  
  `GET https://localhost:8080/places/1`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### PATCH `/places/{objectID}`

Updates the place with the given `objectID` and return it as a `place`.

Examples:  
  `PATCH https://localhost:8080/places/1`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/places/{objectID}`

Deletes the place with the given `objectID`.
  
Examples:  
  `DELETE https://localhost:8080/places/1`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `204 No Content` on success and `error` on failure.
* Codes `204`/`40x`/`50x`

------------------------------------------------------------------------------------

## Tag Route

The **tag route** provides endpoints to retrieve and edit tags and associated entities.
This route uses a `tag` object describing a tag in the FestivalsApp.

**`tag`** object

```json
{
    "tag_id":   "int",
    "tag_name": "string"
}
```

| Field       | Description              |
|-------------|--------------------------|
| `tag_id`    | The id of the tag.       |
| `tag_name`  | The name of the tag.     |

------------------------------------------------------------------------------------

#### GET `/tags`

Gets all tags as a list of `tags`s.

Query Parameter:  
  `name`: Filter result by name  
  `ids` : Filter result by IDs

Examples:  
  `GET https://localhost:8080/tags`  
  `GET https://localhost:8080/tags?name=rock`  
  `GET https://localhost:8080/tags?ids=1,8,56`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `200`/`40x`/`50x`

------------------------------------------------------------------------------------

#### POST `/tags`

Creates a new tag with the provided `tag` and returns it as a `tag`.

Examples:  
  `POST https://localhost:8080/tags`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### GET `/tags/{objectID}`

Gets the tag with the given `objectID` and return it as a `tag`.

Examples:  
  `GET https://localhost:8080/tags/1`

**Authorization**
Requires a valid `Api-Key`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### PATCH `/tags/{objectID}`

Updates the tag with the given `objectID` and return it as a `tag`.

Examples:  
  `PATCH https://localhost:8080/tags/1`  
  `BODY: {OBJECT}`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* `data` or `error` field
* Codes `201`/`40x`/`50x`

------------------------------------------------------------------------------------

#### DELETE `/tags/{objectID}`

Deletes the tag with the given `objectID`.

Examples:  
  `DELETE https://localhost:8080/tags/1`

**Authorization**
Requires a valid `JWT` token with the user role set to `ADMIN` or the owning `CREATOR`.

**Response**

* Returns `204 No Content` on success and `error` on failure.
* Codes `204`/`40x`/`50x`

------------------------------------------------------------------------------------
