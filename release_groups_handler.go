package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
)

func ReleaseGroupsHandler(w traffic.ResponseWriter, r *http.Request) {
  artistGid := r.URL.Query().Get("artist_gid")
  releaseGroups, err := models.FindReleaseGroupsByArtistGid(artistGid)

  if err == nil {
    json.NewEncoder(w).Encode(releaseGroups)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(models.InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    w.WriteHeader(http.StatusInternalServerError)
  }
}
