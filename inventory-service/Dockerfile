FROM golang:1.23.4-alpine

# Install build dependencies
RUN apk add --no-cache git wget

# Install health check tool
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.11/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main .

EXPOSE 50051
CMD ["./main"]