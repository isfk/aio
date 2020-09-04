
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	@echo "proto >>>>>>>"
	./script/protoc.sh
	@echo "Done.\n"

.PHONY: build
build:
	go build --mod=vendor -v -o aio *.go

.PHONY: run_api
run_api:
	./aio api --config=./config/config.yaml

.PHONY: run_article
run_article:
	./aio article --config=./config/config.yaml

.PHONY: run_role
run_role:
	./aio role --config=./config/config.yaml

.PHONY: run_tweet
run_tweet:
	./aio tweet --config=./config/config.yaml

.PHONY: run_user
run_user:
	./aio user --config=./config/config.yaml
