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

![Figure 1: Architecture Overview Highlighted](https://github.com/Festivals-App/festivals-documentation/blob/main/images/architecture/export/architecture_overview_server.svg "Figure 1: Architecture Overview Highlighted")

<hr/>
<p align="center">
  <a href="#development">Development</a> •
  <a href="#deployment">Deployment</a> •
  <a href="#engage">Engage</a>
</p>
<hr/>

## Development

The developement of the FestivalsAPI [(see documentation)](./DOCUMENTATION.md) and the festivals-server is heavily dependend on the [festivals-api-ios](https://github.com/Festivals-App/festivals-api-ios) and the [festivals-database](https://github.com/Festivals-App/festivals-database) which provides the persistent storage to the FestivalsAPI. In my regular development workflow i first mock the needed behaviour in the API client library and then start implementing the changes in the festivals-server and after that in the festivals-database. To test whether the festivals-server is working correctly i'm currently relying on downstream tests of the festivals-api-ios framework.

The general documentation for the Festivals App is in the [festivals-documentation](https://github.com/festivals-app/festivals-documentation) repository. The documentation repository contains architecture information, general deployment documentation, templates and other helpful documents.

### Requirements

- [Golang](https://go.dev/) Version 1.24.1+
- [Visual Studio Code](https://code.visualstudio.com/download) 1.98.2+
  - Plugin recommendations are managed via [workspace recommendations](https://code.visualstudio.com/docs/editor/extension-marketplace#_recommended-extensions).
- [Bash script](https://en.wikipedia.org/wiki/Bash_(Unix_shell)) friendly environment

## Deployment

The Go binaries are able to run without system dependencies so there are not many requirements for the system to run the festivals-server binary,
just follow the [**deployment guide**](./operation/DEPLOYMENT.md) for deploying it inside a virtual machine or the [**local deployment guide**](./operation/local/README.md)
for running it on your macOS developer machine.

## Engage

I welcome every contribution, whether it is a pull request or a fixed typo. The best place to discuss questions and suggestions regarding the festivals-server is the [issues](https://github.com/festivals-app/festivals-server/issues/) section.
More general information and a good starting point if you want to get involved is the [festival-documentation](https://github.com/Festivals-App/festivals-documentation) repository.

The following channels are available for discussions, feedback, and support requests:

| Type                     | Channel                                                |
| ------------------------ | ------------------------------------------------------ |
| **General Discussion**   | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="General Discussion"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/question.svg?style=flat-square"></a> </a>   |
| **Other Requests**    | <a href="mailto:simon@festivalsapp.org" title="Email me"><img src="https://img.shields.io/badge/email-Simon-green?logo=mail.ru&style=flat-square&logoColor=white"></a>   |

## Licensing

Copyright (c) 2017-2025 Simon Gaus. Licensed under the [**GNU Lesser General Public License v3.0**](./LICENSE)
