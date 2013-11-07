package main

import (
  "testing"
  assert "github.com/pilu/miniassert"
)


func TestRecordingsHandler_WithExistingGid(t *testing.T) {
  recorder := newTestRequest("GET", "/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings")
  assert.Equal(t, 200, recorder.Code)
}

func TestRecordingsHandler_WithExistingArtistGidAndGid(t *testing.T) {
  recorder := newTestRequest("GET", "/artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings")
  assert.Equal(t, 200, recorder.Code)
}

func TestRecordingsHandler_WithExistingdGidAndWrongArtistGid(t *testing.T) {
  recorder := newTestRequest("GET", "/artists/00000000-0000-0000-0000-000000000000/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings")

  body := string(recorder.Body.Bytes())
  assert.Equal(t, `{"error":"Not Found"}` + "\n", body)
  assert.Equal(t, 404, recorder.Code)
}


func TestRecordingsHandler_WithGidNotFound(t *testing.T) {
  recorder := newTestRequest("GET", "/releases/00000000-0000-0000-0000-000000000000/recordings")

  body := string(recorder.Body.Bytes())
  assert.Equal(t, `{"error":"Not Found"}` + "\n", body)
  assert.Equal(t, 404, recorder.Code)
}

func TestRecordingsHandler_WithInvalidUUID(t *testing.T) {
  recorder := newTestRequest("GET", "/releases/bad-uuid/recordings")

  body := string(recorder.Body.Bytes())
  assert.Equal(t, "", body)
  assert.Equal(t, 400, recorder.Code)
}
