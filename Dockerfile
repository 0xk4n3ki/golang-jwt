# TO build: sudo docker build --tag jwt-auth-server:v2-latest .
# To run  : sudo docker run -it --rm -v "$(pwd)/.env":/app/.env -p 9000:9000 jwt-auth-server:v2-latest

FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -o server -v .

FROM alpine:3.19

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN addgroup -S app && adduser -S -G app app

WORKDIR /app

COPY --from=builder /app/server .

USER app

EXPOSE 9000

ENTRYPOINT [ "/app/server" ]