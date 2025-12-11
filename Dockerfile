FROM golang:1.22-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/payments-service

FROM debian:bookworm-slim

WORKDIR /app/cmd/server

# Create non-root user
RUN useradd -r -u 10001 appuser

COPY --from=builder /bin/payments-service /usr/local/bin/payments-service
# Default config for container runs; can be overridden by mounting your own.
COPY config.example.yml /app/config.yml

EXPOSE 8080
USER appuser

ENTRYPOINT ["payments-service"]
