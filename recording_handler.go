package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
)

type Recording struct {
  Gid               string           `json:"gid"`
  Name              string           `json:"name"`
  Comment           string           `json:"comment"`
  Length            int              `json:"length"`
  Artists           []*ReleaseArtist `json:"artists"`
}

const FindRecordingByGidQuery = `
  SELECT
    R.gid, R.name, R.comment, R.length, R.artist_credit
  FROM
    recording R
  WHERE
    R.gid = $1 limit 1;`

func FindRecordingByGid(gid string) (*Recording, error) {
  recording := &Recording{}

  if !isValidUUID(gid) {
    return recording, InvalidUUID{ gid }
  }

  var artistCredit int

  row := DB.QueryRow(FindRecordingByGidQuery, gid)
  err := row.Scan(&recording.Gid, &recording.Name, &recording.Comment, &recording.Length, &artistCredit)

  if err != nil {
    return recording, err
  }

  recording.Artists = FindReleaseArtistsByArtistCredit(artistCredit)

  return recording, nil
}

func RecordingHandler(w traffic.ResponseWriter, r *http.Request) {
  gid := r.URL.Query().Get("gid")
  recording, err := FindRecordingByGid(gid)

  if err == nil {
    json.NewEncoder(w).Encode(recording)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    w.WriteHeader(http.StatusInternalServerError)
  }
}
