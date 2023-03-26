FROM golang:1.18-alpine as builder

RUN apk add --update-cache \
    build-base

RUN mkdir -p /var/local/app

COPY . /var/local/app

WORKDIR /var/local/app

RUN go test -v ./...

RUN go build cmd/restapi.go

RUN chmod +x restapi

ENTRYPOINT ["./restapi"]