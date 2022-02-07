run:
	go run ./cmd/simple-server

build:
	go build -ldflags="-s -w" -o ./cmd/simple-server/simple-server.so ./cmd/simple-server