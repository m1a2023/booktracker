FROM golang:1.24 AS builder

RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc libc6-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o app ./cmd/server
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates libsqlite3-0 && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 2020
CMD ["./app"]
