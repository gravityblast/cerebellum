package main

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestReleaseGroupHandler_WithExistingReleaseGroupId(t *testing.T) {
	recorder := newTestRequest("GET", "/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"id":"aa997ea0-2936-40bd-884d-3af8a0e064dc","name":"Random Access Memories","comment":"","firstReleaseDate":"2013-05-17","type":"Album","artists":[{"id":"056e4f3e-d505-4dad-8ec1-d04f521cbb56","name":"Daft Punk"}]}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 200, recorder.Code)
}

func TestReleaseGroupHandler_WithReleaseGroupIdNotFound(t *testing.T) {
	recorder := newTestRequest("GET", "/release-groups/00000000-0000-0000-0000-000000000000")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, `{"error":"release group not found"}`+"\n", body)
	assert.Equal(t, 404, recorder.Code)
}

func TestReleaseGroupHandler_WithInvalidReleaseGroupId(t *testing.T) {
	recorder := newTestRequest("GET", "/release-groups/bad-uuid")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, "", body)
	assert.Equal(t, 400, recorder.Code)
}

func TestReleaseGroupHandler_WithExistingArtistIdAndReleaseGroupId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"id":"aa997ea0-2936-40bd-884d-3af8a0e064dc","name":"Random Access Memories","comment":"","firstReleaseDate":"2013-05-17","type":"Album","artists":[{"id":"056e4f3e-d505-4dad-8ec1-d04f521cbb56","name":"Daft Punk"}]}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 200, recorder.Code)
}

func TestReleaseGroupHandler_WithExistingReleaseGroupIdAndWrongArtistId(t *testing.T) {
	// artist is Guns'n'Roses, release group is Random Access Memories by Daft Punk
	recorder := newTestRequest("GET", "/artists/eeb1195b-f213-4ce1-b28c-8565211f8e43/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"error":"release group not found"}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 404, recorder.Code)
}

func TestReleaseGroupHandler_WithExistingReleaseGroupIdAndNonExistingArtistId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/00000000-0000-0000-0000-000000000000/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"error":"artist not found"}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 404, recorder.Code)
}
