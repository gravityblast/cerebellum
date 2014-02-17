package main

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestReleasesHandler_WithExistingArtistId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases")
	assert.Equal(t, 200, recorder.Code)
}

func TestReleasesHandler_WithArtistIdNotFound(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/00000000-0000-0000-0000-000000000000/releases")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, `{"error":"artist not found"}`+"\n", body)
	assert.Equal(t, 404, recorder.Code)
}

func TestReleasesHandler_WithInvalidUUID(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/bad-uuid/releases")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, "", body)
	assert.Equal(t, 400, recorder.Code)
}

func TestReleasesHandler_WithExistingArtistIdAndReleaseGroupId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc/releases")
	assert.Equal(t, 200, recorder.Code)
}

func TestReleaseHandler_WithExistingArtistIdAndNonExistingReleaseGroupId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups/00000000-0000-0000-0000-000000000000/releases")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"error":"release group not found"}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 404, recorder.Code)
}

func TestReleasesHandler_WithExistingArtistIdAndWrongReleaseGroupId(t *testing.T) {
	// Artist is Guns'n'Roses but release group is "Random Access Memories" by Daft Punk
	recorder := newTestRequest("GET", "/artists/eeb1195b-f213-4ce1-b28c-8565211f8e43/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc/releases")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"error":"release group not found"}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 404, recorder.Code)
}
