#!/bin/sh
redis-server --tls-port 6379 --port 0 \
 --tls-cert-file /usr/local/etc/redis/tls/redis.crt \
 --tls-key-file /usr/local/etc/redis/tls/redis.key \
 --tls-ca-cert-file /usr/local/etc/redis/tls/ca.crt
