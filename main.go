package main

import (
	"github.com/Phisto/eventusserver/config"
	"github.com/Phisto/eventusserver/server"
)

func main() {

	conf := config.GetConfig()

	server := &server.Server{}
	server.Initialize(conf)
	server.Run(":8080")
}
