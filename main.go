package main

import (
	"github.com/Festivals-App/festivals-server/config"
	"github.com/Festivals-App/festivals-server/server"
)

func main() {
	conf := config.GetConfig()
	serverInstance := &server.Server{}
	serverInstance.Initialize(conf)
	serverInstance.Run(":8080")
}
