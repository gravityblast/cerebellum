package main

import (
  "net/http"
  "encoding/json"
)

type RootResponse struct {
  Version string `json:"version"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
  response := RootResponse{ VERSION }
  json.NewEncoder(w).Encode(response)
}
