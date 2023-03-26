FROM golang:1.18-alpine as builder

RUN apk add --update-cache \
    build-base

VOLUME /var/local/app

COPY . /var/local/app

WORKDIR /var/local/app

RUN go test -v ./...

ENTRYPOINT ["go", "run", "./cmd/restapi.go"]
