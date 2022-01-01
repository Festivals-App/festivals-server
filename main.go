package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Festivals-App/festivals-server/server"
	"github.com/Festivals-App/festivals-server/server/config"
	"github.com/Festivals-App/festivals-server/server/status"
)

func main() {

	conf := config.DefaultConfig()
	if len(os.Args) > 1 {
		conf = config.ParseConfig(os.Args[1])
	}

	serverInstance := &server.Server{}
	serverInstance.Initialize(conf)
	serverInstance.Run(conf.ServiceBindAddress + ":" + strconv.Itoa(conf.ServicePort))
}

func PrintInfo() {
	fmt.Println("Version:\t", status.VersionString())
	fmt.Println("Info:\t", status.InfoString())
}
