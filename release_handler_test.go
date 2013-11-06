package main

import (
  "testing"
  assert "github.com/pilu/miniassert"
)


func TestReleaseHandler_WithExistingGid(t *testing.T) {
  recorder := newTestRequest("GET", "/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5")

  body := string(recorder.Body.Bytes())
  expectedBody := `{"gid":"79215cdf-4764-4dee-b0b9-fec1643df7c5","name":"Random Access Memories","comment":"","status":"Official","packaging":"Jewel Case","artists":[{"gid":"056e4f3e-d505-4dad-8ec1-d04f521cbb56","name":"Daft Punk"}]}` + "\n"

  assert.Equal(t, expectedBody, body)
  assert.Equal(t, 200, recorder.Code)
}

func TestReleaseHandler_WithGidNotFound(t *testing.T) {
  recorder := newTestRequest("GET", "/releases/00000000-0000-0000-0000-000000000000")

  body := string(recorder.Body.Bytes())
  assert.Equal(t, `{"error":"Not Found"}` + "\n", body)
  assert.Equal(t, 404, recorder.Code)
}

func TestReleaseHandler_WithInvalidUUID(t *testing.T) {
  recorder := newTestRequest("GET", "/releases/bad-uuid")

  body := string(recorder.Body.Bytes())
  assert.Equal(t, "", body)
  assert.Equal(t, 400, recorder.Code)
}
