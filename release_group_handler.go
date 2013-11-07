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
  gid := r.URL.Query().Get("gid")
  ReleaseGroup, err := releasegroup.ByGid(gid)

  if err == nil {
    json.NewEncoder(w).Encode(ReleaseGroup)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(models.InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    w.WriteHeader(http.StatusInternalServerError)
  }
}
