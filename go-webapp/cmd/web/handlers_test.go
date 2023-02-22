package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {

	// create a slice of anonymous structs containing the name of the test, the URL path to
	var theTests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound},
	}

	var app application
	routes := app.routes()

	// create a test server using the routes
	ts := httptest.NewTLSServer(routes)

	// defer the closing of the test server until the test function has completed
	defer ts.Close()

	pathToTemplates = "./../../templates/"

	// loop through the slice of anonymous structs
	for _, tt := range theTests {
		resp, err := ts.Client().Get(ts.URL + tt.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != tt.expectedStatusCode {
			t.Errorf("%s: expected %d; got %d", tt.name, tt.expectedStatusCode, resp.StatusCode)
		}

	}

}
