FROM alpine:3.22

WORKDIR /app

COPY ./bin/appinfo .

CMD ["/app/appinfo"]