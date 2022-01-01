# Makefile for festivals-server

VERSION=v1.1.1
DATE=$(shell date +"%d-%m-%Y-%H-%M")
REF=1a410efbd13591db07496601ebc7a059dd55cfe9
export

build:
	go build -v -ldflags="-X 'github.com/Festivals-App/festivals-server/server/status.ServerVersion=$(VERSION)' -X 'github.com/Festivals-App/festivals-server/server/status.BuildTime=$(DATE)' -X 'github.com/Festivals-App/festivals-server/server/status.GitRef=$(REF)'" -o festivals-server main.go

install:
	cp festivals-server /usr/local/bin/festivals-server
	cp config_template.toml /etc/festivals-server.conf
	cp operation/service_template.service /etc/systemd/system/festivals-server.service

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