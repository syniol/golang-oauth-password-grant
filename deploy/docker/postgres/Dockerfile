FROM postgres:14-alpine

COPY ./deploy/docker/postgres/migration /docker-entrypoint-initdb.d/

RUN apk add --update-cache openssl \
    && openssl req -x509 -newkey rsa:4096 -nodes -out /var/lib/postgresql/server.crt -keyout /var/lib/postgresql/server.key \
    -days 365 \
    -subj "/C=GB/O=localhost/OU=Domain Control Validated/CN=localhost"

RUN chmod 0600 /var/lib/postgresql/server.key \
    && chmod 0600 /var/lib/postgresql/server.crt

RUN chown postgres:postgres /var/lib/postgresql/server.key \
    && chown postgres:postgres /var/lib/postgresql/server.crt

ENTRYPOINT ["docker-entrypoint.sh"]

CMD [ "-c", "ssl=on" , "-c", "ssl_cert_file=/var/lib/postgresql/server.crt", "-c",\
    "ssl_key_file=/var/lib/postgresql/server.key" ]
