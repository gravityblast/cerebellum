package main

import (
  "testing"
  /* "fmt" */
  assert "github.com/pilu/miniassert"
)


func TestArtistHandler_WithExistingGid(t *testing.T) {
  recorder := newTestRequest("GET", "/artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56")

  body := string(recorder.Body.Bytes())
  expectedBody := `{"gid":"056e4f3e-d505-4dad-8ec1-d04f521cbb56","name":"Daft Punk","sortName":"Daft Punk","comment":"","beginDate":"1992","endDate":"","type":"Group"}` + "\n"

  assert.Equal(t, expectedBody, body)
  assert.Equal(t, 200, recorder.Code)
}

func TestArtistHandler_WithGidNotFound(t *testing.T) {
  recorder := newTestRequest("GET", "/artists/00000000-0000-0000-0000-000000000000")

  body := string(recorder.Body.Bytes())
  assert.Equal(t, `{"error":"Not Found"}` + "\n", body)
  assert.Equal(t, 404, recorder.Code)
}

func TestArtistHandler_WithInvalidUUID(t *testing.T) {
  recorder := newTestRequest("GET", "/artists/bad-uuid")

  body := string(recorder.Body.Bytes())
  assert.Equal(t, "", body)
  assert.Equal(t, 400, recorder.Code)
}
