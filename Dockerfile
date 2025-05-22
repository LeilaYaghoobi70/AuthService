
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /auth-service ./cmd/main.go


# --- Final stage ---
FROM ubuntu:latest

COPY --from=builder /auth-service auth-service
COPY .env .env

EXPOSE 3000

CMD ["/auth-service"]

