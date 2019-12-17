all: fmt test build

fmt:
	@go fmt

test:
	@go test -v

build:
	@go build

run: build
	@./backend

build-linux:
	env GOOS=linux GOARCH=amd64 go build -o backend-linux-amd64

build-windows:
	env GOOS=windows GOARCH=amd64 go build -o backend-windows-amd64
