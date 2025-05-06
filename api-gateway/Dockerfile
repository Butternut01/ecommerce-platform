# Build stage
FROM golang:1.23.4-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git wget

# Set working directory
WORKDIR /build

# Copy only go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:3.19

# Install health check tool
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.11/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

WORKDIR /app
COPY --from=builder /build/main .

EXPOSE 8080
CMD ["./main"]

