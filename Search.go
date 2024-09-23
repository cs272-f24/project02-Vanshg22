package main

import (
	"fmt"
	"strings"

	"github.com/kljensen/snowball" // Import the stemming library to process search terms.
)

// Index is a custom type that maps terms (after stemming and lower-casing) to a map.
// This nested map associates URLs with the frequency of the term at that URL.
type Index map[string]map[string]int

// search queries the inverted index for a given term and returns a map of URLs with their corresponding frequency of the term.
// It returns an error if the term cannot be processed for any reason, such as stemming failures.
func search(index Index, term string) (map[string]int, error) {
	// Convert the search term to lower case to ensure case insensitivity,
	// making the search robust against different capitalizations in the input term.
	lowerCaseTerm := strings.ToLower(term)

	// Stem the lower case term using the snowball library to match the format of the terms stored in the index.
	// This step is crucial because the index is built with stemmed terms, and the search term must be similarly processed
	// to ensure accurate matching.
	stemmedTerm, err := snowball.Stem(lowerCaseTerm, "english", false)
	if err != nil {
		// Return an error if the stemming process fails, encapsulating the error in a more informative message.
		// This helps in diagnosing issues with specific terms or configurations of the stemming library.
		return nil, fmt.Errorf("error stemming the search term: %w", err)
	}

	// Check if the stemmed term is present in the index.
	// If it is, return the map of URLs and frequencies which represent where and how often the term occurs.
	if frequencies, exists := index[stemmedTerm]; exists {
		return frequencies, nil
	}

	// If the term is not found in the index, return an empty map to indicate that there are no occurrences of the term.
	// This empty result set explicitly conveys that the term, although validly processed, does not appear in any indexed content.
	return make(map[string]int), nil
}
