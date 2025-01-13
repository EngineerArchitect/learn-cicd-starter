package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKeyEmptyAuthHeader(t *testing.T) {
	headers := http.Header{}

	apiKey, err := GetAPIKey(headers)

	if apiKey != "" {
		t.Errorf("expected empty apiKey, got %s", apiKey)
	}

	if errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPIKeyMalformedAuthHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer someapikey")

	apiKey, err := GetAPIKey(headers)

	if apiKey != "" {
		t.Errorf("expected empty apiKey, got %s", apiKey)
	}
	if err == nil || err.Error() != "malformed authorization header" {
		t.Errorf("expected error 'malformed authorization header', got %v", err)
	}
}
