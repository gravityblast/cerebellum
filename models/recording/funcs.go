package recording

import (
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/artist"
)

func AllByReleaseGid(releaseGid string) ([]*models.Recording, error) {
  recordings := make([]*models.Recording, 0)

  if !models.IsValidUUID(releaseGid) {
    return recordings, models.InvalidUUID{ releaseGid }
  }

  rows, err := models.DB.Query(queryAllByReleaseGid, releaseGid)
  if err != nil {
    return recordings, err
  }

  for rows.Next() {
    recording := &models.Recording{}
    err := rows.Scan(&recording.Gid, &recording.Name, &recording.Comment, &recording.Length)
    if err != nil {
      return recordings, err
    }

    recordings = append(recordings, recording)
  }

  return recordings, nil
}

func ByGid(gid string) (*models.Recording, error) {
  recording := &models.Recording{}

  if !models.IsValidUUID(gid) {
    return recording, models.InvalidUUID{ gid }
  }

  var artistCredit int

  row := models.DB.QueryRow(queryByGid, gid)
  err := row.Scan(&recording.Gid, &recording.Name, &recording.Comment, &recording.Length, &artistCredit)

  if err != nil {
    return recording, err
  }

  recording.Artists = artist.AllByArtistCredit(artistCredit)

  return recording, nil
}

func ByReleaseGidAndGid(releaseGid, gid string) (*models.Recording, error) {
  recording := &models.Recording{}

  if !models.IsValidUUID(releaseGid) {
    return recording, models.InvalidUUID{ releaseGid }
  }

  if !models.IsValidUUID(gid) {
    return recording, models.InvalidUUID{ gid }
  }

  var artistCredit int

  row := models.DB.QueryRow(queryByReleaseGidAndGid, releaseGid, gid)
  err := row.Scan(&recording.Gid, &recording.Name, &recording.Comment, &recording.Length, &artistCredit)

  if err != nil {
    return recording, err
  }

  recording.Artists = artist.AllByArtistCredit(artistCredit)

  return recording, nil
}
