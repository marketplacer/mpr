APP=mpr

.PHONY: build
build: clean
	go build -o bin/${APP} main.go

build_all: clean
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP)-macos-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP)-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP)-amd64.exe main.go

.PHONY: run
run:
	go run -race main.go

.PHONY: clean
clean:
	go clean
