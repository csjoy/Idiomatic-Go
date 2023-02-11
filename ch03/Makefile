.DEFAULT_GOAL:=build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: lint
	go vet ./...
.PHONY:vet

install: vet
	go install ./...
.PHONY:install

build: install
	go build main.go
.PHONY:build