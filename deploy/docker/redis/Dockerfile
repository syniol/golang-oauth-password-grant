FROM redis:6-alpine

WORKDIR /usr/local/etc/redis

RUN apk add --update-cache openssl \
    && wget https://raw.githubusercontent.com/redis/redis/6.0/utils/gen-test-certs.sh \
    && sh gen-test-certs.sh \
    && rm gen-test-certs.sh \
    && mv tests/tls . \
    && rm -rf tests

VOLUME /usr/local/etc/redis/tls

COPY ./deploy/docker/redis/entrypoint.sh /usr/bin/entrypoint.sh

ENV REDIS_TLS_CERT_FILE=/usr/local/etc/redis/tls/redis.crt
ENV REDIS_TLS_KEY_FILE=/usr/local/etc/redis/tls/redis.key
ENV REDIS_TLS_CA_FILE=/usr/local/etc/redis/tls/ca.crt

WORKDIR /data

#USER myapp
ENTRYPOINT ["entrypoint.sh"]
