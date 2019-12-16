all: fmt test build

fmt:
	@go fmt

test:
	@go test -v

build:
	@go build

build-linux:
	env GOOS=linux GOARCH=amd64 go build -o backend-linux-amd64
