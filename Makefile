.PHONY: wire run build

wire:
	wire gen ./di

run:
	go run cmd/app/main.go

build:
	go build -o myapp cmd/app/main.go
