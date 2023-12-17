package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"gofr.dev/pkg/gofr/request"
)

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(5 * time.Second) // giving time to connect to DB

	tests := []struct {
		desc       string
		method     string
		endpoint   string
		statusCode int
		body       []byte
	}{
		{"get success", http.MethodGet, "patient", http.StatusOK, nil},
		{"create success", http.MethodPost, "patient", http.StatusCreated, []byte(`{"name": "akash","gender": "male","roomNumber": "16","diagnosis": "typhoid"}`)},
		{"get success", http.MethodGet, "patient/7", http.StatusOK, nil},
		{"update success", http.MethodPut, "patient/7", http.StatusOK, []byte(`{"diagnosis": "healthy"}`)},
		{"delete success", http.MethodDelete, "patient/7", http.StatusNoContent, nil},
		{"get unknown endpoint", http.MethodGet, "patient", http.StatusNotFound, nil},
	}

	for i, tc := range tests {
		req, _ := request.NewMock(tc.method, "http://localhost:8000/"+tc.endpoint, bytes.NewBuffer(tc.body))
		c := http.Client{}

		resp, err := c.Do(req)
		if err != nil {
			t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
			continue // move to next test
		}

		if resp.StatusCode != tc.statusCode {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.statusCode, resp.StatusCode, tc.desc)
		}

		_ = resp.Body.Close()
	}
}