FROM golang:1.24 AS builder

WORKDIR /v1

COPY . .

RUN go build -o server

FROM debian:bookworn-slim

COPY --from=builder /v1/server /server

EXPOSE 2020

CMD [ "/server" ]