package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/release"
)

func ReleaseHandler(w traffic.ResponseWriter, r *http.Request) {
  artistGid := r.URL.Query().Get("artist_gid")
  gid       := r.URL.Query().Get("gid")

  var Release *models.Release
  var err     error

  if artistGid == "" {
    Release, err = release.ByGid(gid)
  } else {
    Release, err = release.ByArtistGidAndGid(artistGid, gid)
  }

  if err == nil {
    json.NewEncoder(w).Encode(Release)
  } else if err == sql.ErrNoRows {
    ReleaseNotFoundHandler(w, r)
  } else if _, ok := err.(models.InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    panic(err)
  }
}
