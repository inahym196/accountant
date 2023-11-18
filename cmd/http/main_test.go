package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name  string
		inOut string
	}{
		{inOut: "Hello World"},
		{inOut: "12345"},
		{inOut: "{test: test}"},
	}
	for _, tt := range tests {
		t.Run(tt.inOut, func(t *testing.T) {
			reqBody := bytes.NewBufferString(tt.inOut)
			req := httptest.NewRequest(http.MethodGet, "/", reqBody)
			rec := httptest.NewRecorder()

			handler(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("want %d, but %d", http.StatusOK, rec.Code)
			}
			if rec := rec.Body.String(); rec != tt.inOut {
				t.Errorf("want %s, but %s", tt.inOut, rec)
			}

		})
	}

}
