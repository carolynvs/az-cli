GOPATH ?= $(shell go env GOPATH)

build:
	go build -o az .

install: build
	cp az ${GOPATH}/bin/
