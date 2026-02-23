# Must use the same Go version as the tyk-gateway binary (1.22.6)
FROM golang:1.22.6-alpine AS builder

# CGO is required for -buildmode=plugin
RUN apk add --no-cache gcc musl-dev

WORKDIR /build

# Copy module files first for layer caching
COPY go.mod ./

# Copy plugin source files
COPY add_channel.go transform_response.go ./

# Build both plugins as Go shared objects (.so)
# CGO_ENABLED=1 is required for -buildmode=plugin
RUN CGO_ENABLED=1 go build -buildmode=plugin -o add_channel.so add_channel.go && \
    CGO_ENABLED=1 go build -buildmode=plugin -o transform_response.so transform_response.go

# Export stage - copy .so files out
FROM scratch
COPY --from=builder /build/add_channel.so /plugins/add_channel.so
COPY --from=builder /build/transform_response.so /plugins/transform_response.so
