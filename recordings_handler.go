package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/recording"
)

func RecordingsHandler(w traffic.ResponseWriter, r *http.Request) {
  artistGid   := r.URL.Query().Get("artist_gid")
  releaseGid  := r.URL.Query().Get("release_gid")

  var recordings  []*models.Recording
  var err         error

  if artistGid == "" {
    recordings, err = recording.AllByReleaseGid(releaseGid)
  } else {
    recordings, err = recording.AllByArtistGidAndReleaseGid(artistGid, releaseGid)
  }

  if err == nil {
    json.NewEncoder(w).Encode(recordings)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(models.InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    w.WriteHeader(http.StatusInternalServerError)
  }
}
