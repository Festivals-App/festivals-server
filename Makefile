# Makefile for festivals-server

VERSION=development
DATE=$(shell date +"%d-%m-%Y-%H-%M")
REF=refs/tags/development
DEV_PATH_MAC=$(shell echo ~/Library/Containers/org.festivalsapp.project)
export

build:
	go build -ldflags="-X 'github.com/Festivals-App/festivals-server/server/status.ServerVersion=$(VERSION)' -X 'github.com/Festivals-App/festivals-server/server/status.BuildTime=$(DATE)' -X 'github.com/Festivals-App/festivals-server/server/status.GitRef=$(REF)'" -o festivals-server main.go

install:
	mkdir -p $(DEV_PATH_MAC)/usr/local/bin
	mkdir -p $(DEV_PATH_MAC)/etc
	mkdir -p $(DEV_PATH_MAC)/var/log
	mkdir -p $(DEV_PATH_MAC)/usr/local/festivals-server

	cp operation/local/ca.crt  $(DEV_PATH_MAC)/usr/local/festivals-server/ca.crt
	cp operation/local/server.crt  $(DEV_PATH_MAC)/usr/local/festivals-server/server.crt
	cp operation/local/server.key  $(DEV_PATH_MAC)/usr/local/festivals-server/server.key
	cp operation/local/database-client.crt  $(DEV_PATH_MAC)/usr/local/festivals-server/database-client.crt
	cp operation/local/database-client.key  $(DEV_PATH_MAC)/usr/local/festivals-server/database-client.key
	cp festivals-server $(DEV_PATH_MAC)/usr/local/bin/festivals-server
	chmod +x $(DEV_PATH_MAC)/usr/local/bin/festivals-server
	cp operation/local/config_template_dev.toml $(DEV_PATH_MAC)/etc/festivals-server.conf

run:
	./festivals-server --container="$(DEV_PATH_MAC)"

run-dev:
	$(DEV_PATH_MAC)/usr/local/bin/festivals-identity-server --container="$(DEV_PATH_MAC)" &
	$(DEV_PATH_MAC)/usr/local/bin/festivals-gateway --container="$(DEV_PATH_MAC)" &

stop-dev:
	killall festivals-gateway
	killall festivals-identity-server

clean:
	rm -r festivals-server
