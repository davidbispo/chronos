# syntax=docker/dockerfile:1

# Start from official Go image
FROM golang:1.22.2-alpine

# Set working directory inside container
WORKDIR /app

# Install required packages
RUN apk add --no-cache git curl mysql-client make

# Install migrate CLI
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate /usr/local/bin/

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

EXPOSE 8080

# Build the Go binary
CMD ["go", "run", "."]

# Start app
# CMD ["./chronos"]
# CMD ["sh", "-c", "while true; do sleep 3600; done"]