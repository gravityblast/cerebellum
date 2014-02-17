package recording

import (
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/artist"
)

func AllByReleaseId(releaseId string) ([]*models.Recording, error) {
  recordings := make([]*models.Recording, 0)

  if !models.IsValidUUID(releaseId) {
    return recordings, models.InvalidUUID{ releaseId }
  }

  rows, err := models.DB.Query(queryAllByReleaseId, releaseId)
  if err != nil {
    return recordings, err
  }

  for rows.Next() {
    recording := &models.Recording{}
    err := rows.Scan(&recording.Id, &recording.Name, &recording.Comment, &recording.Length)
    if err != nil {
      return recordings, err
    }

    recordings = append(recordings, recording)
  }

  return recordings, nil
}

func ById(id string) (*models.Recording, error) {
  recording := &models.Recording{}

  if !models.IsValidUUID(id) {
    return recording, models.InvalidUUID{ id }
  }

  var artistCredit int

  row := models.DB.QueryRow(queryById, id)
  err := row.Scan(&recording.Id, &recording.Name, &recording.Comment, &recording.Length, &artistCredit)

  if err != nil {
    return recording, err
  }

  recording.Artists = artist.AllByArtistCredit(artistCredit)

  return recording, nil
}

func ByReleaseIdAndId(releaseId, id string) (*models.Recording, error) {
  recording := &models.Recording{}

  if !models.IsValidUUID(releaseId) {
    return recording, models.InvalidUUID{ releaseId }
  }

  if !models.IsValidUUID(id) {
    return recording, models.InvalidUUID{ id }
  }

  var artistCredit int

  row := models.DB.QueryRow(queryByReleaseIdAndId, releaseId, id)
  err := row.Scan(&recording.Id, &recording.Name, &recording.Comment, &recording.Length, &artistCredit)

  if err != nil {
    return recording, err
  }

  recording.Artists = artist.AllByArtistCredit(artistCredit)

  return recording, nil
}
