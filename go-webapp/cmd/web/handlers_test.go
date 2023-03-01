package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

// TestAppHome tests the home page handler
func TestAppHome(t *testing.T) {

	// create a new request for the home page
	req, _ := http.NewRequest("GET", "/", nil)

	// add a context and session to the request
	req = addContextAndSessionToRequest(req, app)

	// create a new response recorder
	rw := httptest.NewRecorder()

	handler := http.HandlerFunc(app.Home)

	handler.ServeHTTP(rw, req)

	// check the status code is what we expect
	if rw.Code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rw.Code)
	}

	// read the response body
	body, _ := io.ReadAll(rw.Body)

	// check the response body is what we expect
	if !strings.Contains(string(body), "Your request came from") {
		t.Errorf("want %q; got %q", "Welcome to the home page", string(body))
	}

}

// getCtx returns a context with a value added
func getCtx(req *http.Request) context.Context {

	ctx := context.WithValue(req.Context(), contextUserKey, "unknown")

	return ctx
}

// addContextAndSessionToRequest adds a context and session to the request
func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {

	req = req.WithContext(getCtx(req))

	// add the session to the context
	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))

	return req.WithContext(ctx)
}
