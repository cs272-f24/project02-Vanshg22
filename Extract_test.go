package main

import (
	"reflect"
	"testing"
)

func TestExtract(t *testing.T) {
	// Test for simple HTML
	testHTML := `<p>Hello World!</p><a href="https://example.com">Click here</a>`
	expectedWords := []string{"Hello", "World!", "Click", "here"}
	expectedHrefs := []string{"https://example.com"}

	words, hrefs := extractWordsAndHrefs(testHTML)
	if !reflect.DeepEqual(words, expectedWords) {
		t.Errorf("extractWordsAndHrefs words got %+v, want %+v", words, expectedWords)
	}
	if !reflect.DeepEqual(hrefs, expectedHrefs) {
		t.Errorf("extractWordsAndHrefs hrefs got %+v, want %+v", hrefs, expectedHrefs)
	}

	// Test for multiple links
	testHTML = `<p>Welcome</p><a href="https://example.com">Example</a><a href="https://google.com">Google</a>`
	expectedWords = []string{"Welcome", "Example", "Google"}
	expectedHrefs = []string{"https://example.com", "https://google.com"}

	words, hrefs = extractWordsAndHrefs(testHTML)
	if !reflect.DeepEqual(words, expectedWords) {
		t.Errorf("extractWordsAndHrefs words got %+v, want %+v", words, expectedWords)
	}
	if !reflect.DeepEqual(hrefs, expectedHrefs) {
		t.Errorf("extractWordsAndHrefs hrefs got %+v, want %+v", hrefs, expectedHrefs)
	}

	// Test for no links
	testHTML = `<p>This is just a paragraph without any links.</p>`
	expectedWords = []string{"This", "is", "just", "a", "paragraph", "without", "any", "links."}
	expectedHrefs = []string{}

	words, hrefs = extractWordsAndHrefs(testHTML)
	if !reflect.DeepEqual(words, expectedWords) {
		t.Errorf("extractWordsAndHrefs words got %+v, want %+v", words, expectedWords)
	}
	if len(hrefs) != len(expectedHrefs) {
		t.Errorf("extractWordsAndHrefs hrefs got %+v, want %+v", hrefs, expectedHrefs)
	}

	// Test for empty HTML
	testHTML = ``
	expectedWords = []string{}
	expectedHrefs = []string{}

	words, hrefs = extractWordsAndHrefs(testHTML)
	if len(words) != len(expectedWords) {
		t.Errorf("extractWordsAndHrefs words got %+v, want %+v", words, expectedWords)
	}
	if len(hrefs) != len(expectedHrefs) {
		t.Errorf("extractWordsAndHrefs hrefs got %+v, want %+v", hrefs, expectedHrefs)
	}
}
