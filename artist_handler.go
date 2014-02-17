package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/artist"
)

func ArtistHandler(w traffic.ResponseWriter, r *traffic.Request) {
  id := r.URL.Query().Get("id")
  Artist, err := artist.ById(id)

  if err == nil {
    json.NewEncoder(w).Encode(Artist)
  } else if err == sql.ErrNoRows {
    ArtistNotFoundHandler(w, r)
  } else if _, ok := err.(models.InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    panic(err)
  }
}
