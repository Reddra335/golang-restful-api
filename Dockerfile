# syntax=docker/dockerfile:1

# ---------- Build stage ----------
FROM golang:1.26.1-alpine AS builder

WORKDIR /src

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/restful-api .

# ---------- Runtime stage ----------
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata && \
    adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /out/restful-api /app/restful-api

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/restful-api"]