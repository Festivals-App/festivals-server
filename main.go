package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Festivals-App/festivals-gateway/server/heartbeat"
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
	go sendHeartbeat(conf)
	serverInstance.Run(conf.ServiceBindAddress + ":" + strconv.Itoa(conf.ServicePort))
}

/*
type Heartbeat struct {
	Service   string `json:"service"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Available bool   `json:"available"`
}
*/

func sendHeartbeat(conf *config.Config) {
	for {
		timer := time.After(time.Second * 2)
		<-timer
		var beat *heartbeat.Heartbeat = &heartbeat.Heartbeat{Service: "festivals-server", Host: conf.ServiceBindAddress, Port: conf.ServicePort, Available: true}
		heartbeat.SendHeartbeat(conf.LoversEar, beat)
	}
}

func PrintInfo() {
	fmt.Println("Version:\t", status.VersionString())
	fmt.Println("Info:\t", status.InfoString())
}
