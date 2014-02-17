package main

import (
	"fmt"
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestRootHandler(t *testing.T) {
	recorder := newTestRequest("GET", "/")

	body := string(recorder.Body.Bytes())
	expectedBody := fmt.Sprintf(`{"version":"%s"}`+"\n", VERSION)
	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, []string{"application/json; charset=utf-8"}, recorder.HeaderMap["Content-Type"])
}
