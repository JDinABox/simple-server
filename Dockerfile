FROM golang:alpine AS builderGo
RUN apk --no-cache add --upgrade make

WORKDIR /go/src/github.com/JDinABox/simple-server
COPY go.* ./
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build make build

# Docker build
FROM alpine:latest

RUN apk --no-cache -U upgrade \
  && apk --no-cache add --upgrade ca-certificates \
  && wget -O /bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64 \
  && chmod +x /bin/dumb-init


COPY --from=builderGo /go/src/github.com/JDinABox/simple-server/cmd/simple-server/simple-server.so /bin/simple-server
WORKDIR /etc/simple-server/

# Use dumb-init to prevent gofiber prefork from failing as PID 1
ENTRYPOINT ["/bin/dumb-init", "--"]
CMD ["/bin/simple-server"]