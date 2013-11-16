package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/release"
)

func ReleasesHandler(w traffic.ResponseWriter, r *http.Request) {
  artistGid := r.URL.Query().Get("artist_gid")

  var releases []*models.Release
  var err     error

  releases, err = release.AllByArtistGid(artistGid)

  if err == nil {
    json.NewEncoder(w).Encode(releases)
  } else if err == sql.ErrNoRows {
    ReleaseNotFoundHandler(w, r)
  } else if _, ok := err.(models.InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    panic(err)
  }
}
