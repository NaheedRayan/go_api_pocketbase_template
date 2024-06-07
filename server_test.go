package main

import (
	"net/http"
	"testing"
)

// TestPocketbaseServer tests the functionality of the pocketbase server.
func TestPocketbaseServer(t *testing.T) {
    // Make a request to the pocketbase server.
    response, err := http.Get("http://127.0.0.1:8090/api/")
    if err != nil {
        t.Errorf("Failed to make request: %v", err)
        return
    }
    defer response.Body.Close()

    // Check the response status code.
    if response.StatusCode != http.StatusNotFound {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
    }

    // Add more assertions as needed to validate the response body, headers, etc.
}
