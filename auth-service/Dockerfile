# Dockerfile
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY ./auth-service ./auth-service

RUN go mod download
RUN go build -o main ./auth-service/cmd/main.go

# Runtime stage
FROM debian:bookworm-slim

WORKDIR /root/
COPY --from=builder /app/main .

CMD ["./main"]
