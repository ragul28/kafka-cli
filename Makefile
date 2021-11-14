build:
	GOOS=linux GOARCH=amd64 go build  -ldflags="-s -w" -mod=vendor 

install:
	go install
	
run:
	go build && ./kafka-cli

mod:
	go mod tidy
	go mod verify
	go mod vendor
