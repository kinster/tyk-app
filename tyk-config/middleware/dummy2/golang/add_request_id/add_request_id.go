package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// AddRequestID is the pre-request Go plugin for dummy2.
// It generates a unique X-Request-ID and tags the request with the service name.
func AddRequestID(rw http.ResponseWriter, r *http.Request) {
	requestID := fmt.Sprintf("%016x", rand.Int63())
	r.Header.Set("X-Request-ID", requestID)
	r.Header.Set("X-Service", "dummy2")
}

func main() {}
