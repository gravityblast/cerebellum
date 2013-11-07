package main

import (
  "net/http"
  "encoding/json"
  "github.com/pilu/traffic"
)


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

