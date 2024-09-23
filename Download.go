package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// download fetches the content of the given URL and returns it as a string.
// If the status code of the response is not 200 OK, it returns an error.
func download(url string) (string, error) {
	// Make a GET request to the provided URL
	resp, err := http.Get(url)
	if err != nil {
		// If there is an error making the request (e.g., network issues),
		// return the error.
		return "", err
	}
	// Ensure the response body is closed after reading to prevent resource leaks.
	defer resp.Body.Close()

	// Check if the status code is not 200 OK. If it's any other status code
	// (e.g., 404, 500), return an error indicating a non-success status.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Read the response body content.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// If there is an error reading the body, return the error.
		return "", err
	}

	// Convert the body content from bytes to a string and return it.
	return string(body), nil
}
