# EVE Profit Calculator 2.0 - Backend Production Container
# Multi-stage build for optimized production deployment
# Clean Code: Single Responsibility, Optimized Layers, Security Best Practices

# Build Stage
FROM golang:1.23-alpine AS builder

# Build-time variables for maintainability
ARG SDE_DATABASE_URL="https://www.fuzzwork.co.uk/dump/sqlite-latest.sqlite.bz2"
ARG BINARY_NAME="eve-profit2-backend"

# Install build dependencies (alphabetically sorted for consistency)
RUN apk add --no-cache \
    bzip2 \
    ca-certificates \
    gcc \
    git \
    musl-dev \
    sqlite \
    tzdata \
    wget

# Set working directory
WORKDIR /app

# Copy dependency files first (Docker layer optimization)
COPY go.mod go.sum ./

# Download dependencies (cached if go.mod/go.sum unchanged)
RUN go mod download

# Copy source code (excluding large data files for better layer caching)
COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/

# Download SDE database during build (eliminates build context issues)
RUN mkdir -p data && \
    wget -O data/sqlite-latest.sqlite.bz2 "${SDE_DATABASE_URL}" && \
    bzip2 -d data/sqlite-latest.sqlite.bz2 && \
    echo "SDE database downloaded: $(ls -lh data/sqlite-latest.sqlite)"

# Build the application with optimized settings
RUN CGO_ENABLED=1 GOOS=linux go build \
    -a -installsuffix cgo \
    -ldflags="-w -s" \
    -o ${BINARY_NAME} \
    cmd/server/main.go

# Production Stage - Minimal Alpine image
FROM alpine:3.19 AS production

# Runtime variables for maintainability  
ARG BINARY_NAME="eve-profit2-backend"
ARG APP_USER="appuser"
ARG APP_GROUP="appgroup"
ARG APP_UID=1001
ARG APP_GID=1001

# Install runtime dependencies (minimal set, alphabetically sorted)
RUN apk --no-cache add \
    ca-certificates \
    sqlite \
    tzdata && \
    # Create non-root user for security
    addgroup -g ${APP_GID} ${APP_GROUP} && \
    adduser -D -s /bin/sh -u ${APP_UID} -G ${APP_GROUP} ${APP_USER}

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/${BINARY_NAME} ./main

# Copy SDE database from builder stage
COPY --from=builder /app/data ./data

# Set secure ownership and permissions
RUN chown -R ${APP_USER}:${APP_GROUP} /app && \
    chmod +x ./main

# Switch to non-root user for security
USER ${APP_USER}

# Expose application port
EXPOSE 9000

# Health check with optimized settings
HEALTHCHECK --interval=30s \
            --timeout=3s \
            --start-period=5s \
            --retries=3 \
    CMD wget --no-verbose \
             --tries=1 \
             --spider \
             http://localhost:9000/api/v1/health || exit 1

# Run the application
CMD ["./main"]
