package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
)

type Release struct {
  Gid               string                 `json:"gid"`
  Name              string                 `json:"name"`
  Comment           string                 `json:"comment"`
  Status            string                 `json:"status"`
  Packaging         string                 `json:"packaging"`
  Artists           []*ReleaseGroupArtist  `json:"artists"`
}

const FindReleaseByGidQuery = `
  SELECT
    R.gid, R.name, R.comment, R.artist_credit, RS.name, RP.name
  FROM
    release R
  LEFT JOIN release_status RS
    ON R.status = RS.id
  LEFT JOIN release_packaging RP
    ON R.packaging = RP.id
  WHERE
    R.gid = $1 limit 1;`

func FindReleaseByGid(gid string) (*Release, error) {
  release := &Release{}

  if !isValidUUID(gid) {
    return release, InvalidUUID{ gid }
  }

  var status        *sql.NullString
  var packaging     *sql.NullString
  var artistCredit  int

  row := DB.QueryRow(FindReleaseByGidQuery, gid)
  err := row.Scan(&release.Gid, &release.Name, &release.Comment, &artistCredit, &status, &packaging)

  if err != nil {
    return release, err
  }

  if status != nil {
    release.Status = status.String
  }

  if packaging != nil {
    release.Packaging = packaging.String
  }

  release.Artists = FindArtistsByArtistCredit(artistCredit)

  return release, nil
}

func ReleaseHandler(w traffic.ResponseWriter, r *http.Request) {
  gid := r.URL.Query().Get("gid")
  release, err := FindReleaseByGid(gid)

  if err == nil {
    json.NewEncoder(w).Encode(release)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    w.WriteHeader(http.StatusInternalServerError)
  }
}
