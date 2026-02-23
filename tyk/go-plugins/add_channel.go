package main

import (
	"net/http"

	"github.com/TykTechnologies/tyk/ctx"
	"github.com/TykTechnologies/tyk/log"
)

var logger = log.Get()

// AddChannelHeader is the pre-request Go plugin equivalent of add-channel.js.
// It injects X-CHANNEL and X-Forwarded-Prefix headers before the request
// is forwarded to the upstream backend.
func AddChannelHeader(rw http.ResponseWriter, r *http.Request) {
	channel := "Web" // default channel

	// Override from session metadata if present
	session := ctx.GetSession(r)
	if session != nil {
		if val, ok := session.MetaData["channel"]; ok {
			if ch, ok := val.(string); ok && ch != "" {
				channel = ch
			}
		}
	}

	r.Header.Set("X-CHANNEL", channel)
	r.Header.Set("X-Forwarded-Prefix", "/dummy-proxy/v1")

	logger.WithField("channel", channel).Info("AddChannelHeader: injected X-CHANNEL header")
}

func main() {}
