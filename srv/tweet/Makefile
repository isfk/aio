
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	@echo "making proto >>>>>>>"
	./script/protoc.sh
	@echo -e "\n"

.PHONY: build
build: 
	go build --mod=vendor -v -o tweet-service *.go

.PHONY: run
run: 
	./tweet-service

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t tweet-service:latest
