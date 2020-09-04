
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	@echo "making proto >>>>>>>"
	./script/protoc.sh
	@echo -e "\n"

.PHONY: build
build:
	go build --mod=vendor -v -o article-service *.go

.PHONY: run
run:
	./article-service

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t article-service:latest
