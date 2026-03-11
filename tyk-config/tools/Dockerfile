FROM alpine:latest

RUN apk add --no-cache go git

RUN go install github.com/TykTechnologies/tyk-sync@latest
RUN go install github.com/TykTechnologies/tyk-plugin-compiler@latest

ENV PATH=/root/go/bin:$PATH
