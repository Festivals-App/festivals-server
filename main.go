package main

import (
	"github.com/Phisto/eventusserver/config"
	"github.com/Phisto/eventusserver/server"
)

func main() {
	conf := config.GetConfig()
	serverInstance := &server.Server{}
	serverInstance.Initialize(conf)
	serverInstance.Run(":8080")
}
