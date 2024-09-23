package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrawl(t *testing.T) {
	// Define a test for a simple single link crawl
	t.Run("Simple crawl", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("<html><body><a href=\"/Foo\">Foo</a></body></html>"))
		})
		server := httptest.NewServer(handler)
		defer server.Close()

		visited = make(map[string]bool) // Reset the visited map

		crawl(server.URL) // Call the crawl function

		if !visited[server.URL] {
			t.Errorf("Expected %s to be visited", server.URL)
		}
		expectedFooURL := server.URL + "/Foo"
		if !visited[expectedFooURL] {
			t.Errorf("Expected %s to be visited", expectedFooURL)
		}
	})

	// Test for avoiding re-crawling visited URLs
	t.Run("Avoid re-crawl", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer server.Close()

		visited = make(map[string]bool)
		visited[server.URL] = true // Pretend the URL has already been visited

		crawl(server.URL) // Attempt to crawl again

		// Check that the function respects the visited map
		if len(visited) != 1 {
			t.Errorf("URL was re-crawled: %s", server.URL)
		}
	})

	// Test crawling multiple URLs
	t.Run("Multiple URLs", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`<html><body><a href="/Foo">Foo</a><a href="/Bar">Bar</a></body></html>`))
		})
		server := httptest.NewServer(handler)
		defer server.Close()

		visited = make(map[string]bool)
		crawl(server.URL) // Call the crawl function

		if !visited[server.URL] {
			t.Errorf("Expected %s to be visited", server.URL)
		}
		if !visited[server.URL+"/Foo"] {
			t.Errorf("Expected %s to be visited", server.URL+"/Foo")
		}
		if !visited[server.URL+"/Bar"] {
			t.Errorf("Expected %s to be visited", server.URL+"/Bar")
		}
	})
}
