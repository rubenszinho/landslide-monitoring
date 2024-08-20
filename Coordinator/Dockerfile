# Stage 1: Build the Go application
FROM golang:1.18 as builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o coordinator .

# Stage 2: Set up Mosquitto and run the Go application
FROM ubuntu:latest

# Install Mosquitto and other required tools
RUN apt-get update && apt-get install -y \
    mosquitto \
    mosquitto-clients \
    curl

# Copy the Mosquitto configuration
COPY mosquitto.conf /mosquitto/config/mosquitto.conf

# Copy the built Go application from the builder stage
COPY --from=builder /app/coordinator /usr/local/bin/coordinator

# Expose Mosquitto ports
EXPOSE 1883

# Command to start Mosquitto and the Go application
COPY start.sh /start.sh
ENTRYPOINT ["/bin/bash", "/start.sh"]
