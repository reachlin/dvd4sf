.DEFAULT_GOAL := build

build: main.go
	go build -v main.go

