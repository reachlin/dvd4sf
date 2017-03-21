.DEFAULT_GOAL := build

build: main.go driver.go
	go build -v ./...

install: build
	go install ./...

