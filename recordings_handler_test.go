package main

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestRecordingsHandler_WithExistingReleaseId(t *testing.T) {
	recorder := newTestRequest("GET", "/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings")
	assert.Equal(t, 200, recorder.Code)
}

func TestRecordingsHandler_WithReleaseIdNotFound(t *testing.T) {
	recorder := newTestRequest("GET", "/releases/00000000-0000-0000-0000-000000000000/recordings")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, `{"error":"release not found"}`+"\n", body)
	assert.Equal(t, 404, recorder.Code)
}

func TestRecordingsHandler_WithInvalidReleaseId(t *testing.T) {
	recorder := newTestRequest("GET", "/releases/bad-uuid/recordings")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, "", body)
	assert.Equal(t, 400, recorder.Code)
}

func TestRecordingsHandler_WithExistingArtistIdAndReleaseId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings")
	assert.Equal(t, 200, recorder.Code)
}

func TestRecordingsHandler_WithExistingdIdAndNonExsistingArtistId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/00000000-0000-0000-0000-000000000000/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, `{"error":"artist not found"}`+"\n", body)
	assert.Equal(t, 404, recorder.Code)
}

func TestRecordingsHandler_WithExistingdIdAndWrongArtistId(t *testing.T) {
	// Artist is Guns'n'Roses but release is "Random Access Memories" by Daft Punk
	recorder := newTestRequest("GET", "/artists/eeb1195b-f213-4ce1-b28c-8565211f8e43/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, `{"error":"release not found"}`+"\n", body)
	assert.Equal(t, 404, recorder.Code)
}
