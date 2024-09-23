package main

import (
	"strings"

	"golang.org/x/net/html"
)

// extractWordsAndHrefs parses the provided HTML content and extracts both:
// 1. Words from text nodes (i.e., visible text)
// 2. Hrefs from anchor ("a") tags.
//
// It returns two slices: one for the extracted words and another for the hrefs.
func extractWordsAndHrefs(body string) ([]string, []string) {
	// Parse the HTML content from the input body string.
	// TODO: Add support to omit and skip over style tags.
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		// If there's an error parsing the HTML (e.g., invalid HTML structure),
		// return empty slices for both words and hrefs.
		return []string{}, []string{}
	}

	// Initialize two slices to store extracted words and hrefs.
	var words []string
	var hrefs []string

	// Declare a recursive function to traverse the HTML node tree.
	var traverse func(*html.Node)

	// Define the function to process each HTML node recursively.
	traverse = func(n *html.Node) {
		// Check if the node is a text node, which contains visible text.
		if n.Type == html.TextNode {
			// Use strings.Fields to split the text into words and append to the words slice.
			words = append(words, strings.Fields(n.Data)...)
		}

		// Check if the node is an anchor tag ("a") with an href attribute.
		if n.Type == html.ElementNode && n.Data == "a" {
			// Loop through the node's attributes to find the href attribute.
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					// Append the href value to the hrefs slice.
					hrefs = append(hrefs, attr.Val)
					break // Exit the loop once href is found.
				}
			}
		}

		// Recursively traverse all the child nodes of the current node.
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	// Start the recursive traversal from the root document node.
	traverse(doc)

	// Return the words and hrefs slices.
	// If there were no words or hrefs, empty slices are returned.
	return words, hrefs
}

func main() {
	// Intentionally left blank for the autograder or additional code later.
}
