<p align="center">
   <a href="https://github.com/festivals-app/festivals-server/commits/" title="Last Commit"><img src="https://img.shields.io/github/last-commit/festivals-app/festivals-server?style=flat"></a>
   <a href="https://github.com/festivals-app/festivals-server/issues" title="Open Issues"><img src="https://img.shields.io/github/issues/festivals-app/festivals-server?style=flat"></a>
   <a href="https://github.com/festivals-app/festivals-server" title="SLSA Level 2"><img src="https://img.shields.io/badge/SLSA-Level_2-blue"></a>
   <a href="./LICENSE" title="License"><img src="https://img.shields.io/github/license/festivals-app/festivals-server.svg"></a>
</p>

<h1 align="center">
  <br/><br/>
    FestivalsApp Server
  <br/><br/>
</h1>

A lightweight service providing a RESTful API, called **FestivalsAPI**. The FestivalsAPI provides all the necessary data for the FestivalsApp, including festivals and events, as well as functionality to create and manage this data.

![Figure 1: Architecture Overview Highlighted](https://github.com/Festivals-App/festivals-documentation/blob/main/images/architecture/architecture_overview_server.svg "Figure 1: Architecture Overview Highlighted")

The FestivalsApp Server is designed to be accessed through a FestivalsApp Gateway and connects to a FestivalsApp Database, where the data is stored. A client wanting to call the FestivalsAPI needs a valid mTLS certificate 

<hr/>
<p align="center">
  <a href="#development">Development</a> •
  <a href="#deployment">Deployment</a> •
  <a href="#engage">Engage</a>
</p>
<hr/>

## Development

The developement of the FestivalsAPI [(see documentation)](./DOCUMENTATION.md) and the festivals-server is heavily dependend on the [festivals-api-ios](https://github.com/Festivals-App/festivals-api-ios) and the [festivals-database](https://github.com/Festivals-App/festivals-database) which provides the persistent storage to the FestivalsAPI. In my regular development workflow i first mock the needed behaviour in the API client library and then start implementing the changes in the festivals-server and after that in the festivals-database. To test whether the festivals-server is working correctly i'm currently relying on downstream tests of the festivals-api-ios framework.

To find out more about the architecture and technical information see the [ARCHITECTURE](./ARCHITECTURE.md) document. The general documentation for the Festivals App is in the [festivals-documentation](https://github.com/festivals-app/festivals-documentation) repository. The documentation repository contains architecture information, general deployment documentation, templates and other helpful documents.

### Requirements

- [Golang](https://go.dev/) Version 1.23.02
- [Visual Studio Code](https://code.visualstudio.com/download) 1.94.2+
  - Plugin recommendations are managed via [workspace recommendations](https://code.visualstudio.com/docs/editor/extension-marketplace#_recommended-extensions).
- [Bash script](https://en.wikipedia.org/wiki/Bash_(Unix_shell)) friendly environment

## Deployment

The Go binaries are able to run without system dependencies so there are not many requirements for the system to run the festivals-server binary.
The config file needs to be placed at `/etc/festivals-server.conf` or the template config file needs to be present in the directory the binary runs in.

You also need to provide certificates in the right format and location:

- The default path to the root CA certificate is          `/usr/local/festivals-server/ca.crt`
- The default path to the server certificate is           `/usr/local/festivals-server/server.crt`
- The default path to the corresponding key is            `/usr/local/festivals-server/server.key`
- The default path to the database client certificate is  `/usr/local/festivals-server/database-client.crt`
- The default path to the corresponding key is            `/usr/local/festivals-server/database-client.key`

Where the root CA certificate is required to validate incoming requests, the server certificate and key is required to make outgoing connections
and the database client certificate and key is required to make connections to the [festivals-database](https://github.com/Festivals-App/festivals-database) service.
For instructions on how to manage and create the certificates see the [festivals-pki](https://github.com/Festivals-App/festivals-pki) repository.

### VM

The install, update and uninstall scripts should work with any system that uses *systemd* and *firewalld*.
Additionally the scripts will somewhat work under macOS but won't configure the firewall or launch service.

```bash
#Installing
curl -o install.sh https://raw.githubusercontent.com/Festivals-App/festivals-server/main/operation/install.sh
chmod +x install.sh
sudo ./install.sh

#Updating
curl -o update.sh https://raw.githubusercontent.com/Festivals-App/festivals-server/main/operation/update.sh
chmod +x update.sh
sudo ./update.sh

#To see if the server is running use:
sudo systemctl status festivals-server
```

#### Build and run using make

```bash
make build
make run
# Default API Endpoint : http://localhost:10439
```

## Engage

I welcome every contribution, whether it is a pull request or a fixed typo. The best place to discuss questions and suggestions regarding the festivals-server is the [issues](https://github.com/festivals-app/festivals-server/issues/) section. More general information and a good starting point if you want to get involved is the [festival-documentation](https://github.com/Festivals-App/festivals-documentation) repository.

The following channels are available for discussions, feedback, and support requests:

| Type                     | Channel                                                |
| ------------------------ | ------------------------------------------------------ |
| **General Discussion**   | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="General Discussion"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/question.svg?style=flat-square"></a> </a>   |
| **Other Requests**    | <a href="mailto:simon.cay.gaus@gmail.com" title="Email me"><img src="https://img.shields.io/badge/email-Simon-green?logo=mail.ru&style=flat-square&logoColor=white"></a>   |

### Licensing

Copyright (c) 2017-2024 Simon Gaus. Licensed under the [**GNU Lesser General Public License v3.0**](./LICENSE)
