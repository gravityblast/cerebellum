package main

import (
	"encoding/json"
	"github.com/pilu/traffic"
)

type RootResponse struct {
	Version string `json:"version"`
}

func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	response := RootResponse{VERSION}
	json.NewEncoder(w).Encode(response)
}
