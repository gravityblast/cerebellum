package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "fmt"
  assert "github.com/pilu/miniassert"
)

func TestRootHandler(t *testing.T) {
  assert.False(t, isValidUUID("bad UUI"))

  r, _ := http.NewRequest("GET", "", nil)
  w := httptest.NewRecorder()
  RootHandler(w, r)

  body := string(w.Body.Bytes())
  expectedBody := fmt.Sprintf(`{"version":"%s"}`, VERSION)
  assert.Equal(t, expectedBody, body)
}
