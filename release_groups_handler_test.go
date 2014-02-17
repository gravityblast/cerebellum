package main

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestReleaseGroupsHandler_WithExistingArtistId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups.json")
	assert.Equal(t, 200, recorder.Code)
}

func TestReleaseGroupsHandler_WithNonExistingArtistId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/00000000-0000-0000-0000-000000000000/release-groups.json")
	body := string(recorder.Body.Bytes())
	assert.Equal(t, `{"error":"artist not found"}`+"\n", body)
	assert.Equal(t, 404, recorder.Code)
}

func TestReleaseGroupsHandler_WithInvalidArtistId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/00000000/release-groups.json")
	body := string(recorder.Body.Bytes())
	assert.Equal(t, "", body)
	assert.Equal(t, 400, recorder.Code)
}
