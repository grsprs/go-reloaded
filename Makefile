.PHONY: build test clean install

build:
	go build -o bin/go-reloaded ./cmd/go-reloaded

test:
	go test ./tests -v

clean:
	rm -rf bin/

install: build
	cp bin/go-reloaded $(GOPATH)/bin/

release:
	GOOS=windows GOARCH=amd64 go build -o bin/go-reloaded-windows.exe ./cmd/go-reloaded
	GOOS=linux GOARCH=amd64 go build -o bin/go-reloaded-linux ./cmd/go-reloaded
	GOOS=darwin GOARCH=amd64 go build -o bin/go-reloaded-macos ./cmd/go-reloaded