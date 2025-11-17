FROM alpine:3.22

WORKDIR /app

COPY ./bin/applogin .

CMD ["/app/applogin"]