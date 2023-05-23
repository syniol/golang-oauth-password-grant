FROM redis:6-alpine

RUN addgroup -S myapp \
    && adduser -S -u 10000 -g myapp myapp

RUN mkdir -p /usr/local/etc/redis/tls

WORKDIR /usr/local/etc/redis/tls

RUN apk add --update-cache openssl \
    && wget https://raw.githubusercontent.com/redis/redis/6.0/utils/gen-test-certs.sh \
    && sh gen-test-certs.sh \
    && rm gen-test-certs.sh

ENV ALLOW_EMPTY_PASSWORD=false
ENV REDIS_PASSWORD=SyniolIsTheFuture
ENV REDIS_DISABLE_COMMANDS=FLUSHDB,FLUSHALL
ENV REDIS_TLS_CERT_FILE=/usr/local/etc/redis/tls/redis.crt
ENV REDIS_TLS_KEY_FILE=/usr/local/etc/redis/tls/redis.key
ENV REDIS_TLS_CA_FILE=/usr/local/etc/redis/tls/ca.crt
ENV REDIS_TLS_ENABLED=yes
ENV REDIS_TLS_PORT=6379

WORKDIR /data

#USER myapp
