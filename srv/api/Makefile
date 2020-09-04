
GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	go build --mod=vendor -v -o api-web *.go

.PHONY: run
run: 
	./api-web

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t api-web:latest
