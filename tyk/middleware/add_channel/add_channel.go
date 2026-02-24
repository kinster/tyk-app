package main

import (
	"net/http"
)

// AddChannelHeader is the pre-request Go plugin equivalent of add-channel.js.
// It injects X-CHANNEL and X-Forwarded-Prefix headers before the request
// is forwarded to the upstream backend.
// Tyk calls this function by name as defined in the API definition.
func AddChannelHeader(rw http.ResponseWriter, r *http.Request) {
	channel := "Web" // default channel

	// Override from X-Session-Channel if set upstream
	if sessionChannel := r.Header.Get("X-Session-Channel"); sessionChannel != "" {
		channel = sessionChannel
	}

	r.Header.Set("X-CHANNEL", channel)
	r.Header.Set("X-Forwarded-Prefix", "/dummy-proxy/v1")
}

func main() {}
