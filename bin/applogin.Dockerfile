FROM alpine:3.22

WORKDIR /app

COPY ./applogin .

CMD ["./applogin"]