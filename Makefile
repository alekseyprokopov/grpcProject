.PHONY: gen run clear test
test:
	go test -cover -race ./...
run:
	go run main.go
gen:
	protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
clear:
	rm -r pb/*
