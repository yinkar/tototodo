package src

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	// GET request to /health endpoint
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Creating a response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)

	// Serve HTTP to make requests
	handler.ServeHTTP(rr, req)

	// Check status code
	status := rr.Code
	assert.Equal(t, status, http.StatusOK)

	// Check body
	body := rr.Body.String()
	expected := `{"alive": true}`
	assert.Equal(t, body, expected)
}

func TestVersion(t *testing.T) {
	// GET request to /health endpoint
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Creating a response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Version)

	// Serve HTTP to make requests
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("The endpoint returned %v but we want %v as a status code.",
			status,
			http.StatusOK)
	}

	// Check body
	expected := `{"version": "v0.1.0"}`
	if rr.Body.String() != expected {
		t.Errorf("The endpoint returned %v but we want %v as a body.",
			rr.Body.String(), expected)
	}
}
