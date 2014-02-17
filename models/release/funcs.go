package release

import (
  "database/sql"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/artist"
)

func ById(id string) (*models.Release, error) {
  release := &models.Release{}

  if !models.IsValidUUID(id) {
    return release, models.InvalidUUID{ id }
  }

  var status        *sql.NullString
  var packaging     *sql.NullString
  var _type        *sql.NullString
  var artistCredit  int

  row := models.DB.QueryRow(queryById, id)
  err := row.Scan(&release.Id, &release.Name, &release.Comment, &artistCredit, &status, &packaging, &_type)

  if err != nil {
    return release, err
  }

  if status != nil {
    release.Status = status.String
  }

  if packaging != nil {
    release.Packaging = packaging.String
  }

  if _type != nil {
    release.Type = _type.String
  }

  release.Artists = artist.AllByArtistCredit(artistCredit)

  return release, nil
}

func ByArtistIdAndId(artistId, id string) (*models.Release, error) {
  release, err := ById(id)
  if err != nil {
    return release, err
  }

  for _, artist := range release.Artists {
    if artist.Id == artistId {
      return release, nil
    }
  }

  return release, sql.ErrNoRows
}

func AllByArtistId(artistId string) ([]*models.Release, error) {
  releases  := make([]*models.Release, 0)
  rows, err := models.DB.Query(queryAllByArtistId, artistId)
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

    err := rows.Scan(&release.Id, &release.Name, &release.Comment, &status, &_type, &packaging, &dateYear, &dateMonth, &dateDay)
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

func AllByReleaseGroupId(releaseGroupId string) ([]*models.Release, error) {
  releases  := make([]*models.Release, 0)
  rows, err := models.DB.Query(queryAllByReleaseGroupId, releaseGroupId)
  if err != nil {
    return releases, err
  }

  for rows.Next() {
    release := &models.Release{}

    var status    *sql.NullString
    var _type     *sql.NullString
    var packaging *sql.NullString

    err := rows.Scan(&release.Id, &release.Name, &release.Comment, &status, &_type, &packaging)
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

    releases = append(releases, release)
  }

  return releases, nil
}

func Exists(id string) bool {
  if !models.IsValidUUID(id) {
    return false
  }

  var found int

  row := models.DB.QueryRow(queryExists, id)
  err := row.Scan(&found)
  if err != nil {
    return false
  }

  return true
}
