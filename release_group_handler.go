package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/releasegroup"
)

func ReleaseGroupHandler(w traffic.ResponseWriter, r *http.Request) {
  artistGid := r.URL.Query().Get("artist_gid")
  gid       := r.URL.Query().Get("gid")

  var ReleaseGroup *models.ReleaseGroup
  var err          error

  if artistGid != "" {
    ReleaseGroup, err = releasegroup.ByArtistGidAndGid(artistGid, gid)
  } else {
    ReleaseGroup, err = releasegroup.ByGid(gid)
  }

  if err == nil {
    json.NewEncoder(w).Encode(ReleaseGroup)
  } else if err == sql.ErrNoRows {
    ReleaseGroupNotFoundHandler(w, r)
  } else if _, ok := err.(models.InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    panic(err)
  }
}
