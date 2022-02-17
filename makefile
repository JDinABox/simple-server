run:
	go run ./cmd/simple-server

build:
	go build -ldflags="-s -w" -o ./cmd/simple-server/simple-server.so ./cmd/simple-server

install:
	go install -ldflags="-s -w" ./cmd/simple-server

qtc:
	qtc -dir=app/template

docker-build:
	docker build --tag simple-server:latest . 