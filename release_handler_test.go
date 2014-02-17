package main

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestReleaseHandler_WithExistingId(t *testing.T) {
	recorder := newTestRequest("GET", "/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"id":"79215cdf-4764-4dee-b0b9-fec1643df7c5","name":"Random Access Memories","comment":"","status":"Official","type":"Album","packaging":"Jewel Case","artists":[{"id":"056e4f3e-d505-4dad-8ec1-d04f521cbb56","name":"Daft Punk"}]}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 200, recorder.Code)
}

func TestReleaseHandler_WithIdNotFound(t *testing.T) {
	recorder := newTestRequest("GET", "/releases/00000000-0000-0000-0000-000000000000.json")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, `{"error":"release not found"}`+"\n", body)
	assert.Equal(t, 404, recorder.Code)
}

func TestReleaseHandler_WithInvalidUUID(t *testing.T) {
	recorder := newTestRequest("GET", "/releases/bad-uuid.json")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, "", body)
	assert.Equal(t, 400, recorder.Code)
}

func TestReleaseHandler_WithExistingArtistIdAndId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"id":"79215cdf-4764-4dee-b0b9-fec1643df7c5","name":"Random Access Memories","comment":"","status":"Official","type":"Album","packaging":"Jewel Case","artists":[{"id":"056e4f3e-d505-4dad-8ec1-d04f521cbb56","name":"Daft Punk"}]}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 200, recorder.Code)
}

func TestReleaseHandler_WithExistingIdAndNonExistingArtistId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/00000000-0000-0000-0000-000000000000/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"error":"artist not found"}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 404, recorder.Code)
}

func TestReleaseHandler_WithExistingIdAndWrongArtistId(t *testing.T) {
	// Artist is Guns'n'Roses but release is "Random Access Memories" by Daft Punk
	recorder := newTestRequest("GET", "/artists/eeb1195b-f213-4ce1-b28c-8565211f8e43/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"error":"release not found"}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 404, recorder.Code)
}
