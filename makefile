run:
	go run ./cmd/simple-server

build: qtc
	go build -ldflags="-s -w" -o ./cmd/simple-server/simple-server.so ./cmd/simple-server

qtc:
	qtc -dir=app/template