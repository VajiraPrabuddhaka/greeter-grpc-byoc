.PHONY: proto build run client docker-build docker-run clean

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/greeting.proto

build: proto
	go build -o bin/server .
	go build -o bin/client ./client

run: build
	./bin/server

client: build
	./bin/client

docker-build:
	docker build -t greeter-grpc .

docker-run: docker-build
	docker run -p 50051:50051 greeter-grpc

clean:
	rm -rf bin/
	rm -f proto/*.pb.go

deps:
	go mod tidy
	go mod download