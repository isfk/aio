
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	@echo "making proto >>>>>>>"
	./script/protoc.sh
	@echo -e "\n"

.PHONY: build
build:
	go build --mod=vendor -v -o user-service *.go

.PHONY: run
run:
	./user-service

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-service:latest
