FROM redis:6-alpine

RUN addgroup -S myapp \
    && adduser -S -u 10000 -g myapp myapp

USER myapp
