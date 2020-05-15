MODULE=git.gzsunrun.cn/sunruniaas/goctl
VERSION_PKG=$(MODULE)/pkg/version
BINARY=bin/goctl
GO_VERSION=$(shell go version)
BUILD_TIME=$(shell date +%FT%T%z)
GIT_COMMIT=$(shell git rev-parse HEAD)
APP_VERSION=$(shell git rev-parse --abbrev-ref HEAD | grep -v HEAD || git describe --exact-match HEAD 2> /dev/null || git rev-parse HEAD)

LDFLAGS=-ldflags "-w -s -X '$(VERSION_PKG).AppVersion=${APP_VERSION}' -X '$(VERSION_PKG).BuildTime=${BUILD_TIME}' -X '$(VERSION_PKG).GoVersion=${GO_VERSION}' -X '$(VERSION_PKG).GitCommit=${GIT_COMMIT}'"

.PHONY: bin
all: bin

gen:
	go run gen/generate.go

bin:
	go build -o $(BINARY) main.go tpl.go