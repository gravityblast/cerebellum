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
  artistGid       := r.URL.Query().Get("artist_gid")
  releaseGroupGid := r.URL.Query().Get("release_group_gid")

  var releases []*models.Release
  var err     error

  if releaseGroupGid == "" {
    releases, err = release.AllByArtistGid(artistGid)
  } else {
    releases, err = release.AllByReleaseGroupGid(releaseGroupGid)
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
