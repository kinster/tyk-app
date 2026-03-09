package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// EnrichedResponse wraps the upstream response with metadata.
type EnrichedResponse struct {
	Data      interface{} `json:"data"`
	RequestID string      `json:"requestId"`
	Service   string      `json:"service"`
	Enriched  bool        `json:"enriched"`
}

// EnrichResponse is the response Go plugin for dummy2.
// It wraps the upstream response body with service metadata.
func EnrichResponse(rw http.ResponseWriter, res *http.Response, req *http.Request) {
	if res == nil {
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	res.Body.Close()

	var upstream interface{}
	if err := json.Unmarshal(body, &upstream); err != nil {
		res.Body = io.NopCloser(bytes.NewBuffer(body))
		return
	}

	enriched := EnrichedResponse{
		Data:      upstream,
		RequestID: req.Header.Get("X-Request-ID"),
		Service:   "dummy2",
		Enriched:  true,
	}

	newBody, err := json.Marshal(enriched)
	if err != nil {
		res.Body = io.NopCloser(bytes.NewBuffer(body))
		return
	}

	res.Body = io.NopCloser(bytes.NewBuffer(newBody))
	res.Header.Set("Content-Type", "application/json")
	res.Header.Set("Content-Length", fmt.Sprintf("%d", len(newBody)))
	res.ContentLength = int64(len(newBody))
}

func main() {}
