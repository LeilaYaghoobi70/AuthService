
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /docker-gs-ping ./cmd/main.go


# --- Final stage ---
FROM ubuntu:latest

LABEL authors="leilayaghoobi"

COPY --from=builder /docker-gs-ping /docker-gs-ping

EXPOSE 3000

CMD ["/docker-gs-ping"]

