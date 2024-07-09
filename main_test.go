package main

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// 1. Setup
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey expectedApiKey")
	expectedAPIKey := "expectedApiKey"

	// 2. Invocation
	apiKey, err := GetAPIKey(headers)

	// 3. Assertions
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if apiKey != expectedAPIKey {
		t.Errorf("Expected %v, but got %v", expectedAPIKey, apiKey)
	}
}

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Errorf("Expected error, but got none")
	}
}

func TestGetAPIKey_MalformedAuthHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "IncorrectFormat")

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Errorf("Expected error, but got none")
	}
}

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
