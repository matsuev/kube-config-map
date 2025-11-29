FROM alpine:3.22

WORKDIR /app

COPY ./appinfo .

CMD ["./appinfo"]