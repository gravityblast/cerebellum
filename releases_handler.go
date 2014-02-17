package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/release"
)

func ReleasesHandler(w traffic.ResponseWriter, r *traffic.Request) {
  artistId       := r.URL.Query().Get("artist_id")
  releaseGroupId := r.URL.Query().Get("release_group_id")

  var releases []*models.Release
  var err     error

  if releaseGroupId == "" {
    releases, err = release.AllByArtistId(artistId)
  } else {
    releases, err = release.AllByReleaseGroupId(releaseGroupId)
  }

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
