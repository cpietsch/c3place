all: fmt test build

fmt:
	@go fmt ./...

test:
	@GO111MODULE=on go test -v ./...

build:
	@GO111MODULE=on go build

run: build
	@./backend

clean:
	rm -rf backend*

build-all: clean build-darwin build-linux build-windows

build-darwin:
	env GOOS=darwin GOARCH=amd64 go build -o backend-darwin-amd64

build-linux:
	env GOOS=linux GOARCH=amd64 go build -o backend-linux-amd64

build-windows:
	env GOOS=windows GOARCH=amd64 go build -o backend-windows-amd64.exe
