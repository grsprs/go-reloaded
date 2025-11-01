.PHONY: build build-web test clean install run-web

build:
	go build -o bin/go-reloaded ./cmd/go-reloaded

build-web:
	go build -o bin/go-reloaded-web ./cmd/go-reloaded-web

build-all: build build-web

test:
	go test ./tests -v

run-web:
	go run ./cmd/go-reloaded-web

clean:
	rm -rf bin/

install: build-all
	cp bin/go-reloaded $(GOPATH)/bin/
	cp bin/go-reloaded-web $(GOPATH)/bin/

release:
	GOOS=windows GOARCH=amd64 go build -o bin/go-reloaded-windows.exe ./cmd/go-reloaded
	GOOS=linux GOARCH=amd64 go build -o bin/go-reloaded-linux ./cmd/go-reloaded
	GOOS=darwin GOARCH=amd64 go build -o bin/go-reloaded-macos ./cmd/go-reloaded
	GOOS=windows GOARCH=amd64 go build -o bin/go-reloaded-web-windows.exe ./cmd/go-reloaded-web
	GOOS=linux GOARCH=amd64 go build -o bin/go-reloaded-web-linux ./cmd/go-reloaded-web
	GOOS=darwin GOARCH=amd64 go build -o bin/go-reloaded-web-macos ./cmd/go-reloaded-web