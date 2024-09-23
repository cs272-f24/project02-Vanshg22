package main

import (
	"fmt"
)

// visited tracks whether a URL has already been visited to prevent re-crawling the same page.
var visited = make(map[string]bool)

// index is an inverted index where each word is mapped to a dictionary.
// The dictionary maps URLs where the word appears to the frequency of its occurrences.
var index = make(map[string]map[string]int)

// Assume that download, cleanUrls, and extractWordsAndHrefs functions are correctly defined elsewhere in your package.

// crawl is a recursive function that processes a URL: downloads its content, extracts words and URLs,
// and crawls further into linked URLs. It handles each URL once to avoid infinite loops.
func crawl(url string) {
	// Check if the URL has already been processed to avoid cyclic crawling.
	if visited[url] {
		return
	}
	fmt.Printf("Crawling: %s\n", url)

	// Download the content of the page.
	body, err := download(url)
	if err != nil {
		fmt.Printf("Error downloading %s: %v\n", url, err)
		return
	}

	// Mark the URL as visited after a successful download.
	visited[url] = true

	// Extract words and hrefs from the downloaded content.
	words, hrefs := extractWordsAndHrefs(body)

	// Update the index with words extracted from the current URL.
	for _, word := range words {
		if _, exists := index[word]; !exists {
			index[word] = make(map[string]int) // Initialize map if word not yet indexed.
		}
		index[word][url]++
	}

	// Clean the extracted URLs and enqueue them for crawling.
	cleanedURLs := cleanUrls(url, hrefs)
	for _, cleanedURL := range cleanedURLs {
		crawl(cleanedURL)
	}
}
