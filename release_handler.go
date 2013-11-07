package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
)

func ReleaseHandler(w traffic.ResponseWriter, r *http.Request) {
  artistGid := r.URL.Query().Get("artist_gid")
  gid       := r.URL.Query().Get("gid")

  var release *models.Release
  var err     error

  if artistGid == "" {
    release, err = models.FindReleaseByGid(gid)
  } else {
    release, err = models.FindReleaseByArtistGidAndGid(artistGid, gid)
  }

  if err == nil {
    json.NewEncoder(w).Encode(release)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(models.InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    w.WriteHeader(http.StatusInternalServerError)
  }
}
