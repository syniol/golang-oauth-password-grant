FROM postgres:14-alpine

COPY ./deploy/docker/postgres/migration /docker-entrypoint-initdb.d/

#RUN chown -R myapp:myapp /docker-entrypoint-initdb.d/
