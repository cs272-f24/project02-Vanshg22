package main

import (
	"fmt"
	"net/url"
)

// cleanUrls resolves relative hrefs to absolute URLs based on the baseURL.
func cleanUrls(baseURL string, hrefs []string) []string {
	base, err := url.Parse(baseURL)
	if err != nil {
		return []string{} // Return empty slice on error
	}

	cleanedUrls := make([]string, len(hrefs))
	for i, href := range hrefs {
		// Parse the href relative to the base URL
		parsedUrl, err := base.Parse(href)
		if err != nil {
			fmt.Printf("Invalid href: %s\n", href)
			continue
		}
		cleanedUrls[i] = parsedUrl.String() // Convert parsed URL to absolute string
	}
	return cleanedUrls
}
