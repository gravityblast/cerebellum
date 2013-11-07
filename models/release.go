package models

import (
  "database/sql"
)

type Release struct {
  Gid               string           `json:"gid"`
  Name              string           `json:"name"`
  Comment           string           `json:"comment"`
  Status            string           `json:"status"`
  Packaging         string           `json:"packaging"`
  Artists           []*ReleaseArtist `json:"artists"`
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

  release.Artists = FindReleaseArtistsByArtistCredit(artistCredit)

  return release, nil
}

func FindReleaseByArtistGidAndGid(artistGid, gid string) (*Release, error) {
  release, err := FindReleaseByGid(gid)
  if err != nil {
    return release, err
  }

  for _, artist := range release.Artists {
    if artist.Gid == artistGid {
      return release, nil
    }
  }

  return release, sql.ErrNoRows
}

