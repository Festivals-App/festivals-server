<h1 align="center">
    Festivals App Server
</h1>

<p align="center">
   <a href="https://github.com/festivals-app/festivals-server/commits/" title="Last Commit"><img src="https://img.shields.io/github/last-commit/festivals-app/festivals-server?style=flat"></a>
   <a href="https://github.com/festivals-app/festivals-server/issues" title="Open Issues"><img src="https://img.shields.io/github/issues/festivals-app/festivals-server?style=flat"></a>
   <a href="./LICENSE" title="License"><img src="https://img.shields.io/github/license/festivals-app/festivals-server.svg"></a>
</p>

<p align="center">
  <a href="#development">Development</a> •
  <a href="#usage">Usage</a> •
  <a href="#deployment">Deployment</a> •
  <a href="#engage">Engage</a> •
  <a href="#licensing">Licensing</a>
</p>

A live and lightweight go server app providing a simple RESTful API using [go-chi/chi](https://github.com/go-chi/chi) and [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql).

## Development

TBA

### Requirements

-  go 1.11

### Setup development

Install homebrew: https://brew.sh/index_de
Setup local mysql environment: https://tableplus.com/blog/2018/11/how-to-download-mysql-mac.html

## Usage

TBA

### Structure
```
├── server
│   ├── server.go               // Server logic
│   │
│   ├── config
│   │   └── config.go           // Server configuration
│   │
│   ├── database               
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
└── main.go               
```

### Documentation

The FestivalsAPI is documented in detail [here](./DOCUMENTATION.md).

The full documentation for the Festivals App is in the [festivals-documentation](https://github.com/festivals-app/festivals-documentation) repository. 
The documentation repository contains technical documents, architecture information and UI/UX specifications related to this implementation.

## Deployment

The install, update and uninstall scripts should work with any system that uses *systemd* and *firewalld* or *ufw*. 
Additionally the scripts will somewhat work under macOS but won't configure the firewall or launch service.

Installing
```bash
curl -o install.sh https://raw.githubusercontent.com/Festivals-App/festivals-server/master/operation/install.sh
chmod +x install.sh
sudo ./install.sh
```
Updating
```bash
curl -o update.sh https://raw.githubusercontent.com/Festivals-App/festivals-server/master/operation/update.sh
chmod +x update.sh
sudo ./update.sh
```
Uninstalling
```bash
curl -o uninstall.sh https://raw.githubusercontent.com/Festivals-App/festivals-server/master/operation/uninstall.sh
chmod +x uninstall.sh
sudo ./uninstall.sh
```
To see if the server is running use:
```bash
sudo systemctl status festivals-server
```

### Docker

```bash
TBA
```

### Build and Run manually
```bash
cd $GOPATH/src/github.com/Festivals-App/festivals-server
go build main.go
./main

# Default API Endpoint : http://localhost:10439
```

## Engage

TBA

The following channels are available for discussions, feedback, and support requests:

| Type                     | Channel                                                |
| ------------------------ | ------------------------------------------------------ |
| **General Discussion**   | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="General Discussion"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/question.svg?style=flat-square"></a> </a>   |
| **Concept Feedback**    | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="Open Concept Feedback"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/architecture.svg?style=flat-square"></a>  |
| **Other Requests**    | <a href="mailto:phisto05@gmail.com" title="Email Festivals Team"><img src="https://img.shields.io/badge/email-Festivals%20team-green?logo=mail.ru&style=flat-square&logoColor=white"></a>   |

## Licensing

Copyright (c) 2020 Simon Gaus.

Licensed under the **GNU Lesser General Public License v3.0** (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License at https://www.gnu.org/licenses/lgpl-3.0.html.

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the [LICENSE](./LICENSE) for the specific language governing permissions and limitations under the License.