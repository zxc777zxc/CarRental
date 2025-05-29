# Build stage
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./notification-service ./notification-service

RUN go build -o /app/main ./notification-service/cmd/main.go

# Runtime stage
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .

CMD ["./main"]
