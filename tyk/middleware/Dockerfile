# Must use Tyk's official plugin compiler matching the gateway version
# This ensures Go version and ABI compatibility with the gateway binary
FROM tykio/tyk-plugin-compiler:v5.11.0 AS builder

WORKDIR /build

# Copy module files first for layer caching
COPY go.mod ./

# Copy plugin source files
COPY add_channel.go transform_response.go ./

# Build both plugins using the Tyk plugin compiler
# The compiler handles CGO and outputs versioned filenames automatically
RUN go build -buildmode=plugin -o add_channel.so add_channel.go && \
    go build -buildmode=plugin -o transform_response.so transform_response.go

# Export stage - alpine has sh so docker run can copy .so files out
FROM alpine:3.20
COPY --from=builder /build/add_channel.so /plugins/add_channel.so
COPY --from=builder /build/transform_response.so /plugins/transform_response.so
