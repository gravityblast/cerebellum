package main

import (
  "net/http"
  "encoding/json"
  "github.com/pilu/traffic"
)

type RootResponse struct {
  Version string `json:"version"`
}

func RootHandler(w traffic.ResponseWriter, r *http.Request) {
  response := RootResponse{ VERSION }
  json.NewEncoder(w).Encode(response)
}
