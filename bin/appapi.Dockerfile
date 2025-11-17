FROM alpine:3.22

WORKDIR /app

COPY ./appapi .

CMD ["./appapi"]