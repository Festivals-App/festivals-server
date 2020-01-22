# Eventus Server

[![License](https://img.shields.io/github/license/phisto/eventusserver.svg)](https://github.com/Phisto/eventusserver)

Eventus server, a live and lightweight go server app.

## Overview

A simple RESTful API with Go using go-chi/chi and go-sql-driver/mysql.

## Requirements

-  go 1.7

## Installation

```bash
go get github.com/Phisto/eventusserver
```

Before running the API server, you should set the database config with your values in config/config.go

```
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "guest",
			Password: "Password0000!",
			Name:     "eventus",
			Charset:  "utf8",
		},
	}
}
```

# Build and Run
```bash
cd $GOPATH/src/github.com/Phisto/eventusserver
go build main.go
./main

# API Endpoint : http://localhost:8080
```

## Structure
```
├── server
│   ├── server.go               // Server logic
│   │     
│   ├── database               
│   │   ├── CreateDatabase.sql  // Script to create the database
│   │   ├── InsertTestData.sql  // Script to insert some test data.
│   │   ├── mysql.go            // Basic mysql queries (SELECT, INSERT, etc.)
│   │   └── querytools.go       // Some tools to create mysql query statements
│   │
│   ├── handler                
│   │   ├── common.go           // Common response functions
│   │   ├── festival.go         // APIs for the Festival model
│   │   ├── artist.go           // APIs for the Artist model
│   │   ├── location.go         // APIs for the Location model
│   │   ├── event.go            // APIs for the Event model
│   │   ├── image.go            // APIs for the Image model
│   │   ├── link.go             // APIs for the Link model
│   │   ├── place.go            // APIs for the Place model
│   │   └── tag.go              // APIs for the Tag model
│   │
│   └── model
│       └── model.go            // The object models
│
├── config
│   └── config.go               // Server configuration
│
└── main.go               
```


## Eventus API Documentation

By providing API documentation for all Eventus services...

### Response

Requests that are handled gracefully by the server will always return a top level object  
with at least either the`data`or`error`field. The`data`field will always contain an array.  
If the request returns any objects they will be in that array,
```
{
    "data": [
        {OBJECT},
        {OBJECT},
        {OBJECT}
    ]
}
```
otherwise an empty array is returned.
```
{
    "data": []
}
```
If the request specified to include relationships the objects are contained in the`included`field.  
**Included relationships will only work if only one object is returned by the request.**
```
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
```
{
    "error": "An error occured"
}
```
------------------------------------------------------------------------------------
------------------------------------------------------------------------------------

## Festival Objects
A simple object that represents a festival.

```
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
    
* Examples:
            
    `POST https://localhost:8080/festivals`  
    `BODY: {OBJECT}`
    
* Returns 

    * Returns the create festival on success.
    * Codes `201`/`40x`/`50x`
    * `data` or `error` field

------------------------------------------------------------------------------------
#### GET `/festivals/{objectID}`

Get the festival with the given `objectID`.

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
 
 * Examples:
       
    `DELETE https://localhost:8080/festivals/1`
    
 * Returns 
 
     * Returns no object.
     * Codes `204`/`40x`/`50x`
     * `data` or `error` field  
 
 ------------------------------------------------------------------------------------
#### GET `/festivals/{objectID}/{image|links|place|tags|events}`

Get the objects that are described by the`{relationship}`.

 * Examples:
      
    `GET https://localhost:8080/festivals/1/image`  
            
             Note: You need to specify the relationship not the associated object type.
             
 * Returns 
 
     * Returns the objects described by the relationship
     * Codes `20x`/`40x`/`50x`
     * `data` or `error` field  
     
------------------------------------------------------------------------------------
#### POST `/festivals/{objectID}/{image|links|place|tags}/{resourceID}`

Adds the object with the given`{resourceID}`to the relationship for the festival with the given`{objectID}`.

 * Examples:
 
    `POST https://localhost:8080/festivals/1/image/1`   
            note: You need to specify the relationship not the associated object type.

 * Returns 
 
     * Returns no object.
     * Codes `200`/`40x`/`50x`
     * `data` or `error` field 

------------------------------------------------------------------------------------
------------------------------------------------------------------------------------

## Artist Objects
A simple object that represents an artist.

```
{
    "artist_id":            integer,
    "artist_version":       string,
    "artist_name":          string,
    "artist_description":   string
}
```

#### GET `/artists`

Get all artists

#### POST `/artists`

Create a new artist

#### GET `/artists/{objectID}`

Get an artist

#### PATCH `/artists/{objectID}`

Update an artist

#### DELETE `/artists/{objectID}`

Delete an artist

#### GET `/artists/{objectID}/{image|links|place|tags}`

Get the given associated objects

#### POST `/artists/{objectID}/{image|links|tags}/{resourceID}`

Associates the object with the given ID with the artist with the given ID

------------------------------------------------------------------------------------
------------------------------------------------------------------------------------

## Location Objects
A simple object that represents a location.

```
{
    "location_id":              integer,
    "location_version":         string,
    "location_name":            string,
    "location_description":     string,
    "location_accessible":      boolean,
    "location_openair":         boolean
}
```

#### GET `/locations`

Get all locations

#### POST `/locations`

Create a new location

#### GET `/locations/{objectID}`

Get a location

#### PATCH `/locations/{objectID}`

Update a location

#### DELETE `/locations/{objectID}`

Delete a location

#### GET `/locations/{objectID}/{image|links|place}`

Get the given associated objects

#### POST `/locations/{objectID}/{image|links|place}/{resourceID}`

Associates the object with the given ID with the location with the given ID

------------------------------------------------------------------------------------
------------------------------------------------------------------------------------

## Event Objects
A simple object that represents an event.

```
{
    "event_id":             integer,
    "event_version":        string,
    "event_name":           string,
    "event_description":    string,
    "event_start":          integer,
    "event_end":            integer
}
```

#### GET `/events`

Get all events

#### POST `/events`

Create a new event

#### GET `/events/{objectID}`

Get an event

#### PATCH `/events/{objectID}`

Update an event

#### DELETE `/events/{objectID}`

Delete an event

#### GET `/events/{objectID}/{artist|location|festival}`

Get the given associated objects

#### POST `/events/{objectID}/{artist|location}/{resourceID}`

Associates the object with the given ID with the event with the given ID

------------------------------------------------------------------------------------
------------------------------------------------------------------------------------

## Image Objects
A simple object that represents an image.

```
{
    "image_id":         integer,
    "image_hash":       string,
    "image_comment":    string,
    "image_ref":        string
}
```

#### GET `/images`

Get all images

#### POST `/images`

Create a new image

#### GET `/images/{objectID}`

Get an image

#### PATCH `/images/{objectID}`

Update an image

#### DELETE `/images/{objectID}`

Delete an image

------------------------------------------------------------------------------------
------------------------------------------------------------------------------------

## Link Objects
A simple object that represents a link.

```
{
    "link_id":      integer,
    "link_version": string,
    "link_url":     string,
    "link_service": integer
}
```

#### GET `/links`

Get all links

#### GET `/links`

Create a new link

#### GET `/links/{objectID}`

Get a link

#### PATCH `/links/{objectID}`

Update a link

#### DELETE `/links/{objectID}`

Delete a link

------------------------------------------------------------------------------------
------------------------------------------------------------------------------------

## Place Objects
A simple object that represents a place.

```
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

#### GET `/places`

Get all places

#### POST `/places`

Create a new place

#### GET `/places/{objectID}`

Get a place

#### PATCH `/places/{objectID}`

Update a place

#### PATCH `/places/{objectID}`

Delete a place

------------------------------------------------------------------------------------
------------------------------------------------------------------------------------

## Image Objects
A simple object that represents a tag.

```
{
    "tag_id":   integer,
    "tag_name": string
}
```

#### GET `/tags`

Get all tags

#### POST `/tags`

Create a new tag

#### GET `/tags/{objectID}`
* `GET`     : Get a tag

#### PATCH `/tags/{objectID}`

Update a tag

#### DELETE `/tags/{objectID}`

Delete a tag

## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 
