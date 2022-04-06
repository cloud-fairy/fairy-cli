#
# Makefile
#
VERSION =
.PHONY: build compile

default: compile

build:
	go build -o bin/fairy

compile:
	rm -rf bin
	rm -rf releases
	GOOS=freebsd GOARCH=386 go build -o bin/$(VERSION)/fairy-freebsd-386 main.go
	GOOS=freebsd GOARCH=amd64 go build -o bin/$(VERSION)/fairy-freebsd-amd64 main.go
	GOOS=linux GOARCH=386 go build -o bin/$(VERSION)/fairy-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/$(VERSION)/fairy-linux-amd64 main.go
	GOOS=windows GOARCH=386 go build -o bin/$(VERSION)/fairy-windows-386.exe main.go
	GOOS=windows GOARCH=amd64 go build -o bin/$(VERSION)/fairy-windows-amd64.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/$(VERSION)/fairy-darwin-amd64 main.go

update-homebrew:
	sed 's/{{.Version}}/${VERSION}/g' ./homebrew/cloudfairy.rb.template > ./homebrew/cloudfairy.rb

run-c-build:
	go run main.go build ../backend/assembler/

run:
	go run main.go generate component /tmp/test-component

run-interactive:
	go run main.go generate component /tmp/test-component --i