# Makefile for festivals-server

VERSION=development
DATE=$(shell date +"%d-%m-%Y-%H-%M")
REF=refs/tags/development
export

build:
	go build -ldflags="-X 'github.com/Festivals-App/festivals-server/server/status.ServerVersion=$(VERSION)' -X 'github.com/Festivals-App/festivals-server/server/status.BuildTime=$(DATE)' -X 'github.com/Festivals-App/festivals-server/server/status.GitRef=$(REF)'" -o festivals-server main.go

install:
	cp festivals-server /usr/local/bin/festivals-server
	cp config_template.toml /etc/festivals-server.conf
	cp operation/service_template.service /etc/systemd/system/festivals-server.service

update:
	systemctl stop festivals-server
	cp festivals-server /usr/local/bin/festivals-server
	systemctl start festivals-server

uninstall:
	rm /usr/local/bin/festivals-server
	rm /etc/festivals-server.conf
	rm /etc/systemd/system/festivals-server.service

run:
	./festivals-server

stop:
	killall festivals-server

clean:
	rm -r festivals-server
