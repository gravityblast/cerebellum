package main

import (
  "net/http"
  "encoding/json"
  "github.com/pilu/traffic"
)

func ErrorHandler(w traffic.ResponseWriter, r *http.Request, err interface{}) {
  json.NewEncoder(w).Encode(map[string]string{
    "error": "something went wrong",
  })
}

func NotFoundHandler(w traffic.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(map[string]string{
    "error": "not found",
  })
}

func SetDefaultHeaders(w traffic.ResponseWriter, r *http.Request) bool {
  w.Header().Set("Cerebellum-Version", VERSION)
  w.Header().Set("Content-Type", "application/json; charset=utf-8")

  return true
}

func ArtistNotFoundHandler(w traffic.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusNotFound)
  json.NewEncoder(w).Encode(map[string]string{
    "error": "artist not found",
  })
}

func ReleaseGroupNotFoundHandler(w traffic.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusNotFound)
  json.NewEncoder(w).Encode(map[string]string{
    "error": "release group not found",
  })
}

func ReleaseNotFoundHandler(w traffic.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusNotFound)
  json.NewEncoder(w).Encode(map[string]string{
    "error": "release not found",
  })
}
