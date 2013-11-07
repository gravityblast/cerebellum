package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/artist"
  "github.com/pilu/cerebellum/models/releasegroup"
)

func ReleaseGroupsHandler(w traffic.ResponseWriter, r *http.Request) {
  artistGid := r.URL.Query().Get("artist_gid")

  if !artist.Exists(artistGid) {
    w.WriteHeader(http.StatusNotFound)
    return
  }

  ReleaseGroups, err := releasegroup.AllByArtistGid(artistGid)

  if err == nil {
    json.NewEncoder(w).Encode(ReleaseGroups)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(models.InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    w.WriteHeader(http.StatusInternalServerError)
  }
}
