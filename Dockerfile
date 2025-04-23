# Start from an official Go image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Install dependencies (e.g. MySQL client)
RUN apk add --no-cache git mysql-client

# Copy go.mod and go.sum before the rest (for layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy project files
COPY . .

# Build the application
RUN go build -o chronos ./main.go

# Set default command
CMD ["./chronos"]