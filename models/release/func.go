package release

import (
  "database/sql"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/artist"
)

func ByGid(gid string) (*models.Release, error) {
  release := &models.Release{}

  if !models.IsValidUUID(gid) {
    return release, models.InvalidUUID{ gid }
  }

  var status        *sql.NullString
  var packaging     *sql.NullString
  var artistCredit  int

  row := models.DB.QueryRow(queryByGid, gid)
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

  release.Artists = artist.AllByArtistCredit(artistCredit)

  return release, nil
}

func ByArtistGidAndGid(artistGid, gid string) (*models.Release, error) {
  release, err := ByGid(gid)
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

