FROM alpine:3.22

WORKDIR /app

COPY ./bin/appapi .

CMD ["/app/appapi"]