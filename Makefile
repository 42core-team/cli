all:
	go run cmd/core-cli/main.go

build:
	go build -o core-cli cmd/core-cli/main.go