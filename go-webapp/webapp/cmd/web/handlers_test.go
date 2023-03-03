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

// TestAppHomeOld tests the home page handler
func TestAppHomeOld(t *testing.T) {

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

// TestAppHome tests the home page handler using a table driven test
func TestAppHome(t *testing.T) {

	var theTests = []struct {
		name            string
		putInSession    string
		expectedContent string
	}{
		{"first visit", "", "Your request came from"},
		{"second visit", "hello, world!", "hello, world!"},
	}

	for _, tt := range theTests {

		// create a new request for the home page
		req, _ := http.NewRequest("GET", "/", nil)

		// add a context and session to the request
		req = addContextAndSessionToRequest(req, app)

		_ = app.Session.Destroy(req.Context())

		if tt.putInSession != "" {

			// add a value to the session if empty
			app.Session.Put(req.Context(), "test", tt.putInSession)
		}

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
		if !strings.Contains(string(body), tt.expectedContent) {
			t.Errorf("want %q; got %q", tt.expectedContent, string(body))
		}

	}
}

// TestApp_renderWithBadTemplate tests the render function with a bad template
func TestApp_renderWithBadTemplate(t *testing.T) {

	// set template path to a bad path
	pathToTemplates = "./testdata"

	// create a new request for the home page
	req, _ := http.NewRequest("GET", "/", nil)

	// add a context and session to the request
	req = addContextAndSessionToRequest(req, app)

	// create a new response recorder
	rw := httptest.NewRecorder()

	error := app.render(rw, req, "bad.page.gothtml", &TemplateData{})

	if error == nil {
		t.Error("expected an error to be returned")
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

