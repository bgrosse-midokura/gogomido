package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestRootHandler(t *testing.T) {
    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(RootHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusMovedPermanently {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    if location := rr.Result().Header.Get("Location"); location != "/health" {
        t.Errorf("handler returned wrong \"Location\" header: got %v want /health",
            location)
    }
}
