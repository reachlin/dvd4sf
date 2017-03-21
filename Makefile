.DEFAULT_GOAL := build

build: main.go driver.go
	go build -v main.go

install: build
	go install ./...

