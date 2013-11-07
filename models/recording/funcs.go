package recording

import (
  "database/sql"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/artist"
)

func allByReleaseGid(releaseGid string) ([]*models.Recording, int, error) {
  recordings := make([]*models.Recording, 0)

  if !models.IsValidUUID(releaseGid) {
    return recordings, 0, models.InvalidUUID{ releaseGid }
  }

  rows, err := models.DB.Query(queryAllByReleaseGid, releaseGid)
  if err != nil {
    return recordings, 0, err
  }

  var artistCredit int

  for rows.Next() {
    recording := &models.Recording{}
    err := rows.Scan(&recording.Gid, &recording.Name, &recording.Comment, &recording.Length, &artistCredit)
    if err == nil {
      recording.Artists = artist.AllByArtistCredit(artistCredit)
      recordings = append(recordings, recording)
    }
  }

  if len(recordings) == 0 {
    return recordings, artistCredit, sql.ErrNoRows
  } else {
    return recordings, artistCredit, nil
  }
}

func AllByReleaseGid(releaseGid string) ([]*models.Recording, error) {
  recordings, _, err := allByReleaseGid(releaseGid)

  return recordings, err
}

func AllByArtistGidAndReleaseGid(artistGid, releaseGid string) ([]*models.Recording, error) {
  recordings, releaseArtistCredit,err := allByReleaseGid(releaseGid)
  if err != nil {
    return recordings, err
  }

  for _, artist := range artist.AllByArtistCredit(releaseArtistCredit) {
    if artist.Gid == artistGid {
      return recordings, nil
    }
  }

  return recordings, sql.ErrNoRows
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
