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

Before running the API server, you should set the database config with your values on config.go

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
go build
./server

# API Endpoint : http://localhost:8080
```

## Structure
```
├── server
│   ├── server.go               // server logic
│   │     
│   ├── database                // Database logic
│   │   ├── CreateDatabase.sql  // Script to create the databse
│   │   ├── InsertTestData.sql  // Script to insert some test data
│   │   ├── mysql.go            // Basic mysql queries (SELECT, INSERT, etc.)
│   │   └── querytools.go       // Some tools to create mysql query statements
│   │
│   ├── handler                 // API handlers
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
│       └── model.go            // The data models
│
├── config
│   └── config.go               // DB Configuration
│
└── main.go               
```


## Eventus API

#### /festivals
* `GET`     : Get all festivals
* `POST`    : Create a new festival

#### /festivals/{objectID}
* `GET`     : Get a festival
* `PATCH`   : Update a festival
* `DELETE`  : Delete a festival

#### /festivals/{objectID}/{image|links|place|tags}
* `GET`     : Get the given associated objects

#### /festivals/{objectID}/{image|links|place|tags}/{resourceID}
* `POST`     : Associates the object with the given ID with the festival with the given ID

#### /artists
* `GET`     : Get all artists
* `POST`    : Create a new artist

#### /artists/{objectID}
* `GET`     : Get an artist
* `PATCH`   : Update an artist
* `DELETE`  : Delete an artist

#### /artists/{objectID}/{image|links|place|tags}
* `GET`     : Get the given associated objects

#### /artists/{objectID}/{image|links|tags}/{resourceID}
* `POST`     : Associates the object with the given ID with the artist with the given ID

#### /locations
* `GET`     : Get all locations
* `POST`    : Create a new location

#### /locations/{objectID}
* `GET`     : Get a location
* `PATCH`   : Update a location
* `DELETE`  : Delete a location

#### /locations/{objectID}/{image|links|place}
* `GET`     : Get the given associated objects

#### /locations/{objectID}/{image|links|place}/{resourceID}
* `POST`    : Associates the object with the given ID with the location with the given ID

#### /events
* `GET`     : Get all events
* `POST`    : Create a new event

#### /events/{objectID}
* `GET`     : Get an event
* `PATCH`   : Update an event
* `DELETE`  : Delete an event

#### /events/{objectID}/{artist|location|festival}
* `GET`     : Get the given associated objects

#### /events/{objectID}/{artist|location}/{resourceID}
* `POST`    : Associates the object with the given ID with the event with the given ID

#### /images
* `GET`     : Get all images
* `POST`    : Create a new image

#### /images/{objectID}
* `GET`     : Get an image
* `PATCH`   : Update an image
* `DELETE`  : Delete an image

#### /links
* `GET`     : Get all links
* `POST`    : Create a new link

#### /links/{objectID}
* `GET`     : Get a link
* `PATCH`   : Update a link
* `DELETE`  : Delete a link

#### /places
* `GET`     : Get all places
* `POST`    : Create a new place

#### /places/{objectID}
* `GET`     : Get a place
* `PATCH`   : Update a place
* `DELETE`  : Delete a place

#### /tags
* `GET`     : Get all tags
* `POST`    : Create a new tag

#### /tags/{objectID}
* `GET`     : Get a tag
* `PATCH`   : Update a tag
* `DELETE`  : Delete a tag

## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 
