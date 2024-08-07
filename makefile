build:
	@go mod tidy
	@go build -o bin/hjkl main.go

run: build
	@HJKL_PORT=5000 ./bin/hjkl