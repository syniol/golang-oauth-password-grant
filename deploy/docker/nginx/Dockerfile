FROM nginx:1.22-alpine

RUN apk add --update-cache \
    curl \
    build-base \
    ca-certificates \
    openssl

RUN openssl req -x509 -newkey rsa:4096 -nodes -out server.crt -keyout server.key \
    -days 365 \
    -subj "/C=GB/O=localhost/OU=Domain Control Validated/CN=localhost"
RUN openssl rsa -in server.key -text > private.pem
RUN openssl x509 -inform PEM -in server.crt > fullchain.pem

COPY ./deploy/docker/nginx/nginx.conf /etc/nginx/nginx.conf
