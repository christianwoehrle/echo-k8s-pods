package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	// Create a request to pass to our handler
	// We don't have any query parameters for now, so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.Handler(handler("test"))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	if rr.Body.String() != "test" {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), version)
	}
}

func TestPromMetric(t *testing.T) {
	// Create a request to pass to our handler
	// We don't have any query parameters for now, so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.Handler(promhttp.Handler())

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	if !strings.Contains(rr.Body.String(), "http_echo_requests_processed_total") {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), version)
	}
}
