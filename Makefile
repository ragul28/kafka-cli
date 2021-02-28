build:
	GOOS=linux GOARCH=amd64 go build -mod=vendor

install:
	go install
	
run:
	go build && ./kafka-cli

mod:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod vendor
