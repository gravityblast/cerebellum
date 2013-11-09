package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/recording"
)

func RecordingHandler(w traffic.ResponseWriter, r *http.Request) {
  releaseGid := r.URL.Query().Get("release_gid")
  gid        := r.URL.Query().Get("gid")

  var Recording *models.Recording
  var err       error

  if releaseGid != "" {
    Recording, err = recording.ByReleaseGidAndGid(releaseGid, gid)
  } else {
    Recording, err = recording.ByGid(gid)
  }

  if err == nil {
    json.NewEncoder(w).Encode(Recording)
  } else if err == sql.ErrNoRows {
    RecordingNotFoundHandler(w, r)
  } else if _, ok := err.(models.InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    panic(err)
  }
}
