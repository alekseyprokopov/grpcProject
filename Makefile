.PHONY: gen clear test server client
test:
	go test -cover -race ./...
gen:
	protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
clear:
	rm -r pb/*
server:
	go run ./cmd/server/main.go -port 8080
client:
	go run ./cmd/client/main.go -address 0.0.0.0:8080

