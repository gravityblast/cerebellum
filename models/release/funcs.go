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

func AllByArtistGid(artistGid string) ([]*models.Release, error) {
  releases  := make([]*models.Release, 0)
  rows, err := models.DB.Query(queryAllByArtistGid, artistGid)
  if err != nil {
    return releases, err
  }

  for rows.Next() {
    release := &models.Release{}

    var status    *sql.NullString
    var _type     *sql.NullString
    var packaging *sql.NullString
    var dateYear  *sql.NullInt64
    var dateMonth *sql.NullInt64
    var dateDay   *sql.NullInt64

    err := rows.Scan(&release.Gid, &release.Name, &release.Comment, &status, &_type, &packaging, &dateYear, &dateMonth, &dateDay)
    if err != nil {
      return releases, err
    }

    if status != nil {
      release.Status = status.String
    }

    if _type != nil {
      release.Type = _type.String
    }

    if packaging != nil {
      release.Packaging = packaging.String
    }

    date := models.DatesToString(dateYear, dateMonth, dateDay)
    if date != "" {
      release.Date = date
    }

    releases = append(releases, release)
  }

  return releases, nil
}

func Exists(gid string) bool {
  if !models.IsValidUUID(gid) {
    return false
  }

  var found int

  row := models.DB.QueryRow(queryExists, gid)
  err := row.Scan(&found)
  if err != nil {
    return false
  }

  return true
}
