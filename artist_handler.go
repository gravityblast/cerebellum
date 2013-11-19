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
  gid := r.URL.Query().Get("gid")
  Artist, err := artist.ByGid(gid)

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
