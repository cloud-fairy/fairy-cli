build:
	go build -o bin/fairy

run:
	go run main.go generate component /tmp/test-component

compile:
	rm -rf bin
	GOOS=freebsd GOARCH=386 go build -o bin/fairy-freebsd-386 main.go
	GOOS=freebsd GOARCH=amd64 go build -o bin/fairy-freebsd-amd64 main.go
	GOOS=linux GOARCH=386 go build -o bin/fairy-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/fairy-linux-amd64 main.go
	GOOS=windows GOARCH=386 go build -o bin/fairy-windows-386.exe main.go
	GOOS=windows GOARCH=amd64 go build -o bin/fairy-windows-amd64.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/fairy-darwin-amd64 main.go
