BINARY=gptcoder
SOURCES=$(wildcard *.go)

.PHONY: all linux windows darwin

all: linux windows darwin

linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY)-linux $(SOURCES)

windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY)-windows $(SOURCES)

darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY)-darwin $(SOURCES)

darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY)-darwin-arm64 $(SOURCES)

