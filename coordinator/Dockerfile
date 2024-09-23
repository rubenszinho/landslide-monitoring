FROM golang:1.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o coordinator .

FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    mosquitto \
    mosquitto-clients \
    curl

COPY mosquitto.conf /mosquitto/config/mosquitto.conf

COPY --from=builder /app/coordinator /usr/local/bin/coordinator

COPY config.env /app/config.env

WORKDIR /app

EXPOSE 1883

COPY start.sh /start.sh
ENTRYPOINT ["/bin/bash", "/start.sh"]
