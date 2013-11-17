package releasegroup

import (
  "database/sql"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/artist"
)

type Scanner interface {
  Scan(...interface{}) error
}

func ScanRecord(scanner Scanner, releaseGroup *models.ReleaseGroup) error {
  var _type *sql.NullString
  var firstReleaseDateYear   *sql.NullInt64
  var firstReleaseDateMonth  *sql.NullInt64
  var firstReleaseDateDay    *sql.NullInt64

  err := scanner.Scan(&releaseGroup.Gid, &releaseGroup.Name, &releaseGroup.Comment, &releaseGroup.ArtistCredit,
                      &_type, &firstReleaseDateYear, &firstReleaseDateMonth, &firstReleaseDateDay)

  if err != nil {
    return err
  }

  date := models.DatesToString(firstReleaseDateYear, firstReleaseDateMonth, firstReleaseDateDay)
  if date != "" {
    releaseGroup.FirstReleaseDate = date
  }

  if _type != nil {
    releaseGroup.Type = _type.String
  }

  return nil
}

func AllByArtistGid(artistGid string) ([]*models.ReleaseGroup, error) {
  releaseGroups := make([]*models.ReleaseGroup, 0)

  if !models.IsValidUUID(artistGid) {
    return releaseGroups, models.InvalidUUID{ artistGid }
  }

  rows, err := models.DB.Query(queryAllByArtistGid, artistGid)
  if err != nil {
    return releaseGroups, err
  }

  for rows.Next() {
    releaseGroup := &models.ReleaseGroup{}
    err := ScanRecord(rows, releaseGroup)
    if err != nil {
      return releaseGroups, err
    }
    releaseGroups = append(releaseGroups, releaseGroup)
  }

  return releaseGroups, nil
}

func ByGid(gid string) (*models.ReleaseGroup, error) {
  releaseGroup := &models.ReleaseGroup{}

  if !models.IsValidUUID(gid) {
    return releaseGroup, models.InvalidUUID{ gid }
  }

  row := models.DB.QueryRow(queryByGid, gid)
  err := ScanRecord(row, releaseGroup)

  if err == nil {
    releaseGroup.Artists = artist.AllByArtistCredit(releaseGroup.ArtistCredit)
  }

  return releaseGroup, err
}

func ByArtistGidAndGid(artistGid, gid string) (*models.ReleaseGroup, error) {
  releaseGroup := &models.ReleaseGroup{}

  if !models.IsValidUUID(artistGid) {
    return releaseGroup, models.InvalidUUID{ artistGid }
  }

  if !models.IsValidUUID(gid) {
    return releaseGroup, models.InvalidUUID{ gid }
  }

  row := models.DB.QueryRow(queryByArtistGidAndGid, artistGid, gid)
  err := ScanRecord(row, releaseGroup)

  if err == nil {
    releaseGroup.Artists = artist.AllByArtistCredit(releaseGroup.ArtistCredit)
  }

  return releaseGroup, err
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

