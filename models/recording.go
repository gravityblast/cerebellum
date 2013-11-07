package models

import (
  "database/sql"
)

type Recording struct {
  Gid               string           `json:"gid"`
  Name              string           `json:"name"`
  Comment           string           `json:"comment"`
  Length            int              `json:"length"`
  Artists           []*ReleaseArtist `json:"artists"`
}

const FindRecordingsByReleaseGidQuery = `
  SELECT
    REC.gid, REC.name, REC.comment, REC.length, REL.artist_credit
  FROM
    recording REC
  JOIN  track T
    ON T.recording = REC.id
  JOIN medium M
    ON T.medium = M.id
  JOIN release REL
    ON M.release = REL.id
  WHERE
    REL.gid = $1
  ORDER BY M.position, T.position;`

func findRecordingsByReleaseGid(releaseGid string) ([]*Recording, int, error) {
  recordings := make([]*Recording, 0)

  if !isValidUUID(releaseGid) {
    return recordings, 0, InvalidUUID{ releaseGid }
  }

  rows, err := DB.Query(FindRecordingsByReleaseGidQuery, releaseGid)
  if err != nil {
    return recordings, 0, err
  }

  var artistCredit int

  for rows.Next() {
    recording := &Recording{}
    err := rows.Scan(&recording.Gid, &recording.Name, &recording.Comment, &recording.Length, &artistCredit)
    if err == nil {
      recording.Artists = FindReleaseArtistsByArtistCredit(artistCredit)
      recordings = append(recordings, recording)
    }
  }

  if len(recordings) == 0 {
    return recordings, artistCredit, sql.ErrNoRows
  } else {
    return recordings, artistCredit, nil
  }
}

func FindRecordingsByReleaseGid(releaseGid string) ([]*Recording, error) {
  recordings, _, err := findRecordingsByReleaseGid(releaseGid)

  return recordings, err
}

func FindRecordingsByArtistGidAndReleaseGid(artistGid, releaseGid string) ([]*Recording, error) {
  recordings, releaseArtistCredit,err := findRecordingsByReleaseGid(releaseGid)
  if err != nil {
    return recordings, err
  }

  for _, artist := range FindReleaseArtistsByArtistCredit(releaseArtistCredit) {
    if artist.Gid == artistGid {
      return recordings, nil
    }
  }

  return recordings, sql.ErrNoRows
}

const FindRecordingByGidQuery = `
  SELECT
    R.gid, R.name, R.comment, R.length, R.artist_credit
  FROM
    recording R
  WHERE
    R.gid = $1 limit 1;`

func FindRecordingByGid(gid string) (*Recording, error) {
  recording := &Recording{}

  if !isValidUUID(gid) {
    return recording, InvalidUUID{ gid }
  }

  var artistCredit int

  row := DB.QueryRow(FindRecordingByGidQuery, gid)
  err := row.Scan(&recording.Gid, &recording.Name, &recording.Comment, &recording.Length, &artistCredit)

  if err != nil {
    return recording, err
  }

  recording.Artists = FindReleaseArtistsByArtistCredit(artistCredit)

  return recording, nil
}
