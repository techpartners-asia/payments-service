# Multi-stage build for optimized production image
# ===============================================

# Stage 1: BUILD (The Build Environment)
# Using a specific Go version is best practice for reproducible builds.
FROM golang:1.25-alpine AS builder

# Set the path to the application's source root inside the container
# NOTE: The deep path is used here to match your project structure, but
# building directly inside /app is generally cleaner.
WORKDIR /app

# Install necessary build tools (only 'make' and 'git' were typically needed for your earlier file)
# If your project uses CGO, you need gcc and libc-dev, otherwise CGO_ENABLED=0 is preferred.
RUN apk add --no-cache bash make git

# 1. Dependency Management (Efficient Layer Caching)
COPY go.mod go.sum ./
RUN go mod download

# 2. Copy source code and build the gRPC server binary
COPY . .

# Environment variables for static Linux binary compilation
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Build the gRPC service
# -ldflags="-w -s" strips debug symbols for a small, production-ready binary.
# The binary is output to the root /app/ for easy copying.
RUN --mount=type=cache,target="/root/.cache/go-build" go build -a -gcflags='-N -l' -installsuffix cgo -o main main.go
# ===============================================
# Stage 2: RUN (The Minimal Runtime Image)
# ===============================================
FROM alpine:latest  

# Install essential runtime dependencies (TZData and Certificates)
RUN apk update && apk add --no-cache tzdata ca-certificates procps

# Set application work directory
WORKDIR /home

# Set Timezone
ENV TZ Asia/Ulaanbaatar

# Copy stripped binary from the builder stage
COPY --from=builder /app/config.yml ./config.yml
COPY --from=builder /app/main ./


# Expose the gRPC port
EXPOSE 50051

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD pgrep -f payments-service-grpc || exit 1

# Default to gRPC server (ENTRYPOINT is good for execution)
ENTRYPOINT ["/home/main"]