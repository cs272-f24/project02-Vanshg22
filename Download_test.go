// TestDownload tests the download function using a mock HTTP server.
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownload(t *testing.T) {
	// Table-driven tests for various scenarios
	tests := []struct {
		name       string
		handler    http.HandlerFunc
		expected   string
		shouldFail bool
	}{
		{
			name: "Successful download",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("<html><body>Hello, World!</body></html>"))
			},
			expected:   "<html><body>Hello, World!</body></html>",
			shouldFail: false,
		},
		{
			name: "404 Not Found",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
			shouldFail: true,
		},
		{
			name: "Empty response",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(""))
			},
			expected:   "",
			shouldFail: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(tc.handler)
			defer server.Close()

			content, err := download(server.URL)
			if tc.shouldFail {
				if err == nil {
					t.Fatalf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("Expected no error, got %v", err)
				}
				if content != tc.expected {
					t.Errorf("Expected %s, got %s", tc.expected, content)
				}
			}
		})
	}

	// Test for invalid URL separately as it doesn't require a mock server
	t.Run("Invalid URL", func(t *testing.T) {
		_, err := download("://invalid-url")
		if err == nil {
			t.Fatalf("Expected an error due to invalid URL, but got none")
		}
	})
}
