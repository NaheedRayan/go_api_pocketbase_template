package main

import (
	"go_api_pocketbase_template/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHello tests the Hello handler.
func TestHello(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	routes.Hello(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	// expected := "Hello from the GOLANG and MUX\n"
	expected := "Hello world\n"
	body := rec.Body.String()
	if body != expected {
		t.Errorf("expected body to be %q; got %q", expected, body)
	}
}


