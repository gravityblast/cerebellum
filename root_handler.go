package main

import (
  "net/http"
  "fmt"
  "encoding/json"
)

type RootResponse struct {
  Version string `json:"version"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
  response := RootResponse{ VERSION }
  responseBytes, _ := json.Marshal(response)
  fmt.Fprint(w, string(responseBytes))
}
