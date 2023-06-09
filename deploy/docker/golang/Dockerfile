FROM golang:1.18-alpine as builder

RUN apk add --update-cache \
    build-base \
    ca-certificates \
    openssl

RUN mkdir -p /var/local/app

COPY . /var/local/app

WORKDIR /var/local/app

RUN go test -v ./...

ENV CGO_ENABLED=0

RUN go build \
    -installsuffix 'static' \
    cmd/restapi.go

RUN openssl req -x509 -newkey rsa:4096 -nodes -out server.crt -keyout server.key \
    -days 365 \
    -subj "/C=GB/O=localhost/OU=Domain Control Validated/CN=localhost"

RUN addgroup -S myapp \
    && adduser -S -u 10000 -g myapp myapp

RUN chown myapp:myapp server.crt  \
    && chown myapp:myapp server.key



FROM scratch

COPY --from=builder /var/local/app/restapi .
COPY --from=builder /var/local/app/server.crt .
COPY --from=builder /var/local/app/server.key .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /etc/passwd /etc/passwd
USER myapp

ENTRYPOINT ["/restapi"]