package main

import (
	"testing"
)

// TestSearch function tests various scenarios using a table-driven approach

func TestSearch(t *testing.T) {
	index := make(Index)
	// Assuming 'romeo' is stemmed and lowercased during indexing
	index["romeo"] = map[string]int{
		"https://cs272-f24.github.io/tests/rnj/sceneI_30.0.html": 2,
		"https://cs272-f24.github.io/tests/rnj/sceneI_30.1.html": 22,
	}

	t.Run("Successful Search", func(t *testing.T) {
		results, _ := search(index, "romeo")
		if len(results) != 2 { // Expecting URLs for 'romeo'
			t.Errorf("Expected 2 results, got %d", len(results))
		}
	})

	t.Run("Term Not Found", func(t *testing.T) {
		results, _ := search(index, "mercutio")
		if len(results) != 0 {
			t.Errorf("Expected 0 results for non-existent term, got %d", len(results))
		}
	})

	t.Run("Case Insensitivity", func(t *testing.T) {
		results, _ := search(index, "RoMeO")
		if len(results) != 2 {
			t.Errorf("Expected 2 results for case insensitivity test, got %d", len(results))
		}
		expectedCount := 22
		if count, exists := results["https://cs272-f24.github.io/tests/rnj/sceneI_30.1.html"]; !exists || count != expectedCount {
			t.Errorf("Expected %d occurrences of 'RoMeO', got %d", expectedCount, count)
		}
	})

	t.Run("Empty Index", func(t *testing.T) {
		emptyIndex := make(Index)
		results, _ := search(emptyIndex, "romeo")
		if len(results) != 0 {
			t.Errorf("Expected 0 results for empty index, got %d", len(results))
		}
	})
}
