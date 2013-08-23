package main

import (
  "testing"
  "fmt"
  assert "github.com/pilu/miniassert"
)

func TestRootHandler(t *testing.T) {
  recorder := newTestRequest("GET", "/")

  body := string(recorder.Body.Bytes())
  expectedBody := fmt.Sprintf(`{"version":"%s"}`, VERSION)
  assert.Equal(t, expectedBody, body)
  assert.Equal(t, 200, recorder.Code)
  assert.Equal(t, []string{"application/json"}, recorder.HeaderMap["Content-Type"])
}
