# Must use the same Go version as the tyk-gateway binary (1.22.6)
# Using Debian-based image (not alpine) as it has gcc/ld pre-installed
# and avoids Alpine aarch64 linker path issues with CGO
FROM golang:1.22.6 AS builder

WORKDIR /build

# Copy module files first for layer caching
COPY go.mod ./

# Copy plugin source files
COPY add_channel.go transform_response.go ./

# Build both plugins as Go shared objects (.so)
# CGO_ENABLED=1 is required for -buildmode=plugin
RUN CGO_ENABLED=1 go build -buildmode=plugin -o add_channel.so add_channel.go && \
    CGO_ENABLED=1 go build -buildmode=plugin -o transform_response.so transform_response.go

# Export stage - alpine has sh so docker run can copy .so files out
FROM alpine:3.20
COPY --from=builder /build/add_channel.so /plugins/add_channel.so
COPY --from=builder /build/transform_response.so /plugins/transform_response.so
