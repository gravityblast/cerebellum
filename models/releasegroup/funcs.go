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

  err := scanner.Scan(&releaseGroup.Id, &releaseGroup.Name, &releaseGroup.Comment, &releaseGroup.ArtistCredit,
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

func AllByArtistId(artistId string) ([]*models.ReleaseGroup, error) {
  releaseGroups := make([]*models.ReleaseGroup, 0)

  if !models.IsValidUUID(artistId) {
    return releaseGroups, models.InvalidUUID{ artistId }
  }

  rows, err := models.DB.Query(queryAllByArtistId, artistId)
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

func ById(id string) (*models.ReleaseGroup, error) {
  releaseGroup := &models.ReleaseGroup{}

  if !models.IsValidUUID(id) {
    return releaseGroup, models.InvalidUUID{ id }
  }

  row := models.DB.QueryRow(queryById, id)
  err := ScanRecord(row, releaseGroup)

  if err == nil {
    releaseGroup.Artists = artist.AllByArtistCredit(releaseGroup.ArtistCredit)
  }

  return releaseGroup, err
}

func ByArtistIdAndId(artistId, id string) (*models.ReleaseGroup, error) {
  releaseGroup := &models.ReleaseGroup{}

  if !models.IsValidUUID(artistId) {
    return releaseGroup, models.InvalidUUID{ artistId }
  }

  if !models.IsValidUUID(id) {
    return releaseGroup, models.InvalidUUID{ id }
  }

  row := models.DB.QueryRow(queryByArtistIdAndId, artistId, id)
  err := ScanRecord(row, releaseGroup)

  if err == nil {
    releaseGroup.Artists = artist.AllByArtistCredit(releaseGroup.ArtistCredit)
  }

  return releaseGroup, err
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

