package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// BackendResponse matches the shape of dummy-backend's /hello response.
type BackendResponse struct {
	Message         string `json:"message"`
	ChannelReceived string `json:"channelReceived"`
	PrefixReceived  string `json:"prefixReceived"`
}

// TransformedResponse is the outbound shape after transformation.
type TransformedResponse struct {
	Status      string `json:"status"`
	Channel     string `json:"channel"`
	Prefix      string `json:"prefix"`
	Transformed bool   `json:"transformed"`
}

// TransformResponse is the response Go plugin.
// It renames fields from the backend response before returning to the client.
// Tyk calls this function by name as defined in the API definition.
func TransformResponse(rw http.ResponseWriter, res *http.Response, req *http.Request) {
	if res == nil {
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	res.Body.Close()

	var upstream BackendResponse
	if err := json.Unmarshal(body, &upstream); err != nil {
		// Pass through unchanged if body can't be parsed
		res.Body = io.NopCloser(bytes.NewBuffer(body))
		return
	}

	transformed := TransformedResponse{
		Status:      upstream.Message,
		Channel:     upstream.ChannelReceived,
		Prefix:      upstream.PrefixReceived,
		Transformed: true,
	}

	newBody, err := json.Marshal(transformed)
	if err != nil {
		res.Body = io.NopCloser(bytes.NewBuffer(body))
		return
	}

	res.Body = io.NopCloser(bytes.NewBuffer(newBody))
	res.Header.Set("Content-Type", "application/json")
	res.Header.Set("Content-Length", fmt.Sprintf("%d", len(newBody)))
	res.ContentLength = int64(len(newBody))
}
