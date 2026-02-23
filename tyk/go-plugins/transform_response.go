package main

import (
	"bytes"
	"encoding/json"
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

// TransformResponse is the response Go plugin equivalent of the transform_response
// Go template. It renames fields and adds transformed:true.
func TransformResponse(rw http.ResponseWriter, res *http.Response, req *http.Request) {
	if res == nil {
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.WithError(err).Error("TransformResponse: failed to read response body")
		return
	}
	res.Body.Close()

	var upstream BackendResponse
	if err := json.Unmarshal(body, &upstream); err != nil {
		logger.WithError(err).Error("TransformResponse: failed to parse response body")
		// Pass through unchanged if we can't parse
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
		logger.WithError(err).Error("TransformResponse: failed to marshal transformed body")
		res.Body = io.NopCloser(bytes.NewBuffer(body))
		return
	}

	res.Body = io.NopCloser(bytes.NewBuffer(newBody))
	res.Header.Set("Content-Type", "application/json")
	res.ContentLength = int64(len(newBody))

	logger.Info("TransformResponse: response body transformed successfully")
}
