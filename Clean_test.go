package main

import (
	"reflect"
	"testing"
)

func TestCleanHref(t *testing.T) {
	tests := []struct {
		name     string   // name of the test case
		baseURL  string   // the base URL for resolving relative links
		hrefs    []string // input slice of hrefs
		expected []string // expected result after cleaning URLs
	}{
		{
			name:     "Only Absolute URLs",
			baseURL:  "https://example.org/",
			hrefs:    []string{"http://google.com", "https://example.org/absolute"},
			expected: []string{"http://google.com", "https://example.org/absolute"},
		},
		{
			name:     "Mixed Relative and Absolute URLs",
			baseURL:  "https://example.org/",
			hrefs:    []string{"/relative", "http://google.com", "/another-relative"},
			expected: []string{"https://example.org/relative", "http://google.com", "https://example.org/another-relative"},
		},
		{
			name:     "Empty Hrefs",
			baseURL:  "https://example.org/",
			hrefs:    []string{},
			expected: []string{},
		},
		{
			name:     "Invalid Base URL",
			baseURL:  "://invalid-url",
			hrefs:    []string{"/foo", "/bar"},
			expected: []string{},
		},
	}

	// Iterate over each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := cleanUrls(tc.baseURL, tc.hrefs)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Test %s failed: cleanUrls got %+v, want %+v", tc.name, result, tc.expected)
			}
		})
	}
}
