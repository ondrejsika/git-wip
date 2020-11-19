build:
	go build

install:
	go install

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o dist/git-wip-darwin-amd64

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o dist/git-wip-linux-amd64
