package models

import (
  "fmt"
  "database/sql"
)

type ReleaseArtist struct {
  Gid   string `json:"gid"`
  Name  string `json:"name"`
}

type ReleaseGroup struct {
  Gid               string           `json:"gid"`
  Name              string           `json:"name"`
  Comment           string           `json:"comment"`
  FirstReleaseDate  string           `json:"firstReleaseDate"`
  Type              string           `json:"type"`
  Artists           []*ReleaseArtist `json:"artists"`
}

const FindReleaseGroupsByArtistGidQuery = `
  SELECT
    RG.gid, RG.name, RG.comment, RG.artist_credit, RGPT.name,
    RGM.first_release_date_year, RGM.first_release_date_month, RGM.first_release_date_day
  FROM
    release_group RG
  JOIN artist_credit_name ACN
    ON RG.artist_credit = ACN.artist_credit
  JOIN artist A
    ON ACN.artist = A.id
  LEFT JOIN release_group_primary_type RGPT
    ON RG.type = RGPT.id
  LEFT JOIN release_group_meta RGM
    ON RG.id = RGM.id
  WHERE
    A.gid = $1;`

func FindReleaseGroupsByArtistGid(artistGid string) ([]*ReleaseGroup, error) {
  releaseGroups := make([]*ReleaseGroup, 0)

  if !isValidUUID(artistGid) {
    return releaseGroups, InvalidUUID{ artistGid }
  }

  rows, err := DB.Query(FindReleaseGroupsByArtistGidQuery, artistGid)
  if err != nil {
    return releaseGroups, err
  }

  for rows.Next() {
    releaseGroup := &ReleaseGroup{}
    var _type *sql.NullString
    var artistCredit int

    var firstReleaseDateYear   *sql.NullInt64
    var firstReleaseDateMonth  *sql.NullInt64
    var firstReleaseDateDay    *sql.NullInt64

    err := rows.Scan(&releaseGroup.Gid, &releaseGroup.Name, &releaseGroup.Comment, &artistCredit, &_type,
                    &firstReleaseDateYear, &firstReleaseDateMonth, &firstReleaseDateDay)

    if err != nil {
      return releaseGroups, err
    }

    if firstReleaseDateYear != nil {
      date := fmt.Sprintf("%d", firstReleaseDateYear.Int64)

      if firstReleaseDateMonth != nil {
        date = fmt.Sprintf("%s-%02d", date, firstReleaseDateMonth.Int64)
      }

      if firstReleaseDateDay != nil {
        date = fmt.Sprintf("%s-%02d", date, firstReleaseDateDay.Int64)
      }

      releaseGroup.FirstReleaseDate = date
    }

    if _type != nil {
      releaseGroup.Type = _type.String
    }

    /* releaseGroup.Artists = FindReleaseArtistsByArtistCredit(artistCredit) */
    releaseGroups = append(releaseGroups, releaseGroup)
  }

  return releaseGroups, nil
}

const FindReleaseGroupByGidQuery = `
  SELECT
    RG.gid, RG.name, RG.comment, RG.artist_credit, RGPT.name,
    RGM.first_release_date_year, RGM.first_release_date_month, RGM.first_release_date_day
  FROM
    release_group RG
  LEFT JOIN release_group_primary_type RGPT
    ON RG.type = RGPT.id
  LEFT JOIN release_group_meta RGM
    ON RG.id = RGM.id
  WHERE
    RG.gid = $1 limit 1;`

func FindReleaseGroupByGid(gid string) (*ReleaseGroup, error) {
  releaseGroup := &ReleaseGroup{}

  if !isValidUUID(gid) {
    return releaseGroup, InvalidUUID{ gid }
  }

  var _type *sql.NullString
  var artistCredit int

  var firstReleaseDateYear   *sql.NullInt64
  var firstReleaseDateMonth  *sql.NullInt64
  var firstReleaseDateDay    *sql.NullInt64

  row := DB.QueryRow(FindReleaseGroupByGidQuery, gid)
  err := row.Scan(&releaseGroup.Gid, &releaseGroup.Name, &releaseGroup.Comment, &artistCredit, &_type,
                  &firstReleaseDateYear, &firstReleaseDateMonth, &firstReleaseDateDay)

  if err != nil {
    return releaseGroup, err
  }

  if firstReleaseDateYear != nil {
    date := fmt.Sprintf("%d", firstReleaseDateYear.Int64)

    if firstReleaseDateMonth != nil {
      date = fmt.Sprintf("%s-%02d", date, firstReleaseDateMonth.Int64)
    }

    if firstReleaseDateDay != nil {
      date = fmt.Sprintf("%s-%02d", date, firstReleaseDateDay.Int64)
    }

    releaseGroup.FirstReleaseDate = date
  }

  if _type != nil {
    releaseGroup.Type = _type.String
  }

  releaseGroup.Artists = FindReleaseArtistsByArtistCredit(artistCredit)

  return releaseGroup, nil
}
