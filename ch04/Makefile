.DEFAULT_GOAL:=build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: lint
	go vet ./...
	shadow ./...
.PHONY:vet

build: vet
	go build main.go
.PHONY:build