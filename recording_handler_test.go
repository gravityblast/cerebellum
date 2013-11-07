package main

import (
  "testing"
  assert "github.com/pilu/miniassert"
)


func TestRecordingHandler_WithExistingGid(t *testing.T) {
  recorder := newTestRequest("GET", "/recordings/833f00e1-781f-4edd-90e4-e52712618862")

  body := string(recorder.Body.Bytes())
  expectedBody := `{"gid":"833f00e1-781f-4edd-90e4-e52712618862","name":"Get Lucky","comment":"","length":367000,"artists":[{"gid":"056e4f3e-d505-4dad-8ec1-d04f521cbb56","name":"Daft Punk"},{"gid":"149f91ef-1287-46da-9a8e-87fee02f1471","name":"Pharrell Williams"}]}` + "\n"

  assert.Equal(t, expectedBody, body)
  assert.Equal(t, 200, recorder.Code)
}

func TestRecordingHandler_WithGidNotFound(t *testing.T) {
  recorder := newTestRequest("GET", "/recordings/00000000-0000-0000-0000-000000000000")

  body := string(recorder.Body.Bytes())
  assert.Equal(t, `{"error":"not found"}` + "\n", body)
  assert.Equal(t, 404, recorder.Code)
}

func TestRecordingHandler_WithInvalidUUID(t *testing.T) {
  recorder := newTestRequest("GET", "/recordings/bad-uuid")

  body := string(recorder.Body.Bytes())
  assert.Equal(t, "", body)
  assert.Equal(t, 400, recorder.Code)
}
