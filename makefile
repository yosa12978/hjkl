.DEFAULT_GOAL = build

build:
	@go mod tidy
	@swag init -g main.go
	@go build -o bin/hjkl main.go

run: build
	@HJKL_PORT=5000 ./bin/hjkl