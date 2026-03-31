
FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd

# ---------- Run Stage ----------
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/app .
COPY --from=builder /app/config ./config

EXPOSE 8080

# Run binary
CMD ["./app"]