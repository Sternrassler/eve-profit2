# EVE Profit Calculator 2.0 - Backend Production Container
# Multi-stage build for optimized production deployment

# Build Stage
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache ca-certificates gcc git musl-dev sqlite tzdata

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code (excluding large data files)
COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/

# Copy data directory separately to avoid cache issues
COPY data/ data/

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Production Stage
FROM alpine:3.19

# Install runtime dependencies and create user
RUN apk --no-cache add ca-certificates sqlite tzdata && \
    addgroup -g 1001 appgroup && \
    adduser -D -s /bin/sh -u 1001 -G appgroup appuser

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main .

# Copy configuration and data files
COPY --from=builder /app/data ./data

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 9000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:9000/api/v1/health || exit 1

# Run the application
CMD ["./main"]
