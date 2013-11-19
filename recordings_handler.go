package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/recording"
)

func RecordingsHandler(w traffic.ResponseWriter, r *traffic.Request) {
  releaseGid  := r.URL.Query().Get("release_gid")

  var recordings  []*models.Recording
  var err         error

  recordings, err = recording.AllByReleaseGid(releaseGid)

  if err == nil {
    json.NewEncoder(w).Encode(recordings)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(models.InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    panic(err)
  }
}
