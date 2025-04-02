# Build Stage
FROM golang:1.24-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache git gcc g++

WORKDIR /app

# Copy dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api/main.go

# Runtime Stage
FROM alpine:3.18

# Add required CA certificates (optional for HTTPS)
RUN apk add --no-cache ca-certificates

# Create a working directory
WORKDIR /app

# Copy the built binary
COPY --from=builder /app/app /app/app

# Set the entrypoint
ENTRYPOINT ["/app/app"]