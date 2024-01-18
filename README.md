<p align="center">
   <a href="https://github.com/festivals-app/festivals-server/commits/" title="Last Commit"><img src="https://img.shields.io/github/last-commit/festivals-app/festivals-server?style=flat"></a>
   <a href="https://github.com/festivals-app/festivals-server/issues" title="Open Issues"><img src="https://img.shields.io/github/issues/festivals-app/festivals-server?style=flat"></a>
   <a href="./LICENSE" title="License"><img src="https://img.shields.io/github/license/festivals-app/festivals-server.svg"></a>
</p>

<h1 align="center">
  <br/><br/>
    Festivals App Server
  <br/><br/>
</h1>

A lightweight go server app providing a RESTful API, called FestivalsAPI. The FestivalsAPI exposes all data functions needed by the FestivalsApp.

![Figure 1: Architecture Overview Highlighted](https://github.com/Festivals-App/festivals-documentation/blob/main/images/architecture/overview_server.png "Figure 1: Architecture Overview Highlighted")

<hr/>
<p align="center">
  <a href="#development">Development</a> •
  <a href="#deployment">Deployment</a> •
  <a href="#festivalsapi">FestivalsAPI</a> •
  <a href="#architecture">Architecture</a> •
  <a href="#engage">Engage</a> •
  <a href="#licensing">Licensing</a>
</p>
<hr/>

## Development

using [go-chi/chi](https://github.com/go-chi/chi) and [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

The developement of the [FestivalsAPI](./DOCUMENTATION.md) and the festivals-server is heavily dependend on the [festivals-api-ios](https://github.com/Festivals-App/festivals-api-ios) as in my regular development workflow i first mock the needed behaviour in the API client library. After the behaviour works the way i need it to, i start implementing the changes in the festivals-server and after that in the [festivals-database](https://github.com/Festivals-App/festivals-database).

To test whether the festivals-server is working correctly i'm currently relying on downstream tests of the [API framework](https://github.com/Festivals-App/festivals-api-ios).

### Requirements

- [Golang](https://go.dev/) Version 1.20+
- [Visual Studio Code](https://code.visualstudio.com/download) 1.84.0+
    * Plugin recommendations are managed via [workspace recommendations](https://code.visualstudio.com/docs/editor/extension-marketplace#_recommended-extensions).
- [Bash script](https://en.wikipedia.org/wiki/Bash_(Unix_shell)) friendly environment

## Deployment

Running the festivals-server is pretty easy because Go binaries are able to run without system dependencies 
on the target for which they are compiled. The only dependency is that the festivals-server expects either a config file at `/etc/festivals-server.conf`,
the environment variables set or the template config file present in the directory it runs from.

### Build and Run manually

```bash
cd /path/to/repository/festivals-server
make build
(make install)
make run

# Default API Endpoint : http://localhost:10439
```

### VM deployment

The install, update and uninstall scripts should work with any system that uses *systemd* and *firewalld*.
Additionally the scripts will somewhat work under macOS but won't configure the firewall or launch service.

Installing
```bash
curl -o install.sh https://raw.githubusercontent.com/Festivals-App/festivals-server/main/operation/install.sh
chmod +x install.sh
sudo ./install.sh
```
Updating
```bash
curl -o update.sh https://raw.githubusercontent.com/Festivals-App/festivals-server/main/operation/update.sh
chmod +x update.sh
sudo ./update.sh
```
To see if the server is running use:
```bash
sudo sudo systemctl status festivals-server
```

## FestivalsAPI

The FestivalsAPI is documented in detail [here](./DOCUMENTATION.md).

## Architecture

The festivals-server is tightly coupled with the [festivals-database](https://github.com/Festivals-App/festivals-database) which provides the persistent storage to the FestivalsAPI. To find out more about architecture and technical information see the [ARCHITECTURE](./ARCHITECTURE.md) document.

The general documentation for the Festivals App is in the [festivals-documentation](https://github.com/festivals-app/festivals-documentation) repository. 
The documentation repository contains architecture information, general deployment documentation, templates and other helpful documents.

## Engage

I welcome every contribution, whether it is a pull request or a fixed typo. The best place to discuss questions and suggestions regarding the festivals-server is the [issues](https://github.com/festivals-app/festivals-server/issues/) section. More general information and a good starting point if you want to get involved is the [festival-documentation](https://github.com/Festivals-App/festivals-documentation) repository.

The following channels are available for discussions, feedback, and support requests:

| Type                     | Channel                                                |
| ------------------------ | ------------------------------------------------------ |
| **General Discussion**   | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="General Discussion"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/question.svg?style=flat-square"></a> </a>   |
| **Other Requests**    | <a href="mailto:simon.cay.gaus@gmail.com" title="Email me"><img src="https://img.shields.io/badge/email-Simon-green?logo=mail.ru&style=flat-square&logoColor=white"></a>   |

## Licensing

Copyright (c) 2017-2023 Simon Gaus.

Licensed under the **GNU Lesser General Public License v3.0** (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License at https://www.gnu.org/licenses/lgpl-3.0.html.

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the [LICENSE](./LICENSE) for the specific language governing permissions and limitations under the License.