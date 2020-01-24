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

## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 
