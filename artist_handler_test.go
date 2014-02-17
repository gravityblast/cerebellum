package main

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestArtistHandler_WithExistingId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"id":"056e4f3e-d505-4dad-8ec1-d04f521cbb56","name":"Daft Punk","sortName":"Daft Punk","comment":"","beginDate":"1992","endDate":"","type":"Group"}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 200, recorder.Code)
}

func TestArtistHandler_WithIdNotFound(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/00000000-0000-0000-0000-000000000000.json")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, `{"error":"artist not found"}`+"\n", body)
	assert.Equal(t, 404, recorder.Code)
}

func TestArtistHandler_WithInvalidUUID(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/bad-uuid.json")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, "", body)
	assert.Equal(t, 400, recorder.Code)
}
