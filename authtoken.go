// Package auth provides functions for extracting a user Auth token from a
// request and associating it with a Context.
package authtoken

import (
	"errors"
	"net/http"
	"strings"
)

const (
	BEARER_SCHEMA string = "Bearer "
)

// FromRequest extracts the auth token from req.
func FromRequest(req *http.Request) (string, error) {
	// Grab the raw Authoirzation header
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header required")
	}

	// Confirm the request is sending Basic Authentication credentials.
	if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
		return "", errors.New("Authorization requires Bearer scheme")
	}

	return authHeader[len(BEARER_SCHEMA):], nil
}
