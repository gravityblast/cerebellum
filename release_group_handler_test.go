package main

import (
  "testing"
  assert "github.com/pilu/miniassert"
)


func TestReleaseGroupHandler_WithExistingGid(t *testing.T) {
  recorder := newTestRequest("GET", "/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc")

  body := string(recorder.Body.Bytes())
  expectedBody := `{"gid":"aa997ea0-2936-40bd-884d-3af8a0e064dc","name":"Random Access Memories","comment":"","firstReleaseDate":"2013-05-17","type":"Album"}` + "\n"

  assert.Equal(t, expectedBody, body)
  assert.Equal(t, 200, recorder.Code)
}

func TestReleaseGroupHandler_WithGidNotFound(t *testing.T) {
  recorder := newTestRequest("GET", "/release-groups/00000000-0000-0000-0000-000000000000")

  body := string(recorder.Body.Bytes())
  assert.Equal(t, `{"error":"Not Found"}` + "\n", body)
  assert.Equal(t, 404, recorder.Code)
}

func TestReleaseGroupHandler_WithInvalidUUID(t *testing.T) {
  recorder := newTestRequest("GET", "/release-groups/bad-uuid")

  body := string(recorder.Body.Bytes())
  assert.Equal(t, "", body)
  assert.Equal(t, 400, recorder.Code)
}
