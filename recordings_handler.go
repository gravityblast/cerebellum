package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
)

const FindRecordingsByReleaseGidQuery = `
  SELECT
    REC.gid, REC.name, REC.comment, REC.length, REL.artist_credit
  FROM
    recording REC
  JOIN  track T
    ON T.recording = REC.id
  JOIN medium M
    ON T.medium = M.id
  JOIN release REL
    ON M.release = REL.id
  WHERE
    REL.gid = $1
  ORDER BY M.position, T.position;`

func findRecordingsByReleaseGid(releaseGid string) ([]*Recording, int, error) {
  recordings := make([]*Recording, 0)

  if !isValidUUID(releaseGid) {
    return recordings, 0, InvalidUUID{ releaseGid }
  }

  rows, err := DB.Query(FindRecordingsByReleaseGidQuery, releaseGid)
  if err != nil {
    return recordings, 0, err
  }

  var artistCredit int

  for rows.Next() {
    recording := &Recording{}
    err := rows.Scan(&recording.Gid, &recording.Name, &recording.Comment, &recording.Length, &artistCredit)
    if err == nil {
      recording.Artists = FindReleaseArtistsByArtistCredit(artistCredit)
      recordings = append(recordings, recording)
    }
  }

  if len(recordings) == 0 {
    return recordings, artistCredit, sql.ErrNoRows
  } else {
    return recordings, artistCredit, nil
  }
}

func FindRecordingsByReleaseGid(releaseGid string) ([]*Recording, error) {
  recordings, _, err := findRecordingsByReleaseGid(releaseGid)

  return recordings, err
}

func FindRecordingsByArtistGidAndReleaseGid(artistGid, releaseGid string) ([]*Recording, error) {
  recordings, releaseArtistCredit,err := findRecordingsByReleaseGid(releaseGid)
  if err != nil {
    return recordings, err
  }

  for _, artist := range FindReleaseArtistsByArtistCredit(releaseArtistCredit) {
    if artist.Gid == artistGid {
      return recordings, nil
    }
  }

  return recordings, sql.ErrNoRows
}

func RecordingsHandler(w traffic.ResponseWriter, r *http.Request) {
  artistGid   := r.URL.Query().Get("artist_gid")
  releaseGid  := r.URL.Query().Get("release_gid")

  var recordings  []*Recording
  var err         error

  if artistGid == "" {
    recordings, err = FindRecordingsByReleaseGid(releaseGid)
  } else {
    recordings, err = FindRecordingsByArtistGidAndReleaseGid(artistGid, releaseGid)
  }

  if err == nil {
    json.NewEncoder(w).Encode(recordings)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    w.WriteHeader(http.StatusInternalServerError)
  }
}
