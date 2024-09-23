FROM golang:1.23.1-alpine3.20 AS build

COPY . .

RUN go build -ldflags "-s -w" -o /server

FROM alpine:3.20

WORKDIR /app

COPY --from=build /server server

CMD [ "/app/server" ]
