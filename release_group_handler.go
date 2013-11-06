package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
)

type ReleaseGroupArtist struct {
  Gid   string `json:"gid"`
  Name  string `json:"name"`
}

type ReleaseGroup struct {
  Gid               string                 `json:"gid"`
  Name              string                 `json:"name"`
  Comment           string                 `json:"comment"`
  FirstReleaseDate  string                 `json:"firstReleaseDate"`
  Type              string                 `json:"type"`
  Artists           []*ReleaseGroupArtist  `json:"artists"`
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


const FindArtistsByArtistCreditQuery = `
  SELECT
    A.gid, A.name
  FROM
    artist_credit_name ACN
  JOIN artist A
    on A.id = ACN.artist
  WHERE
    ACN.artist_credit = $1;`

func FindArtistsByArtistCredit(artistCredit int) []*ReleaseGroupArtist {
  artists := make([]*ReleaseGroupArtist, 0)

  rows, err := DB.Query(FindArtistsByArtistCreditQuery, artistCredit)
  if err != nil {
    return artists
  }

  for rows.Next() {
    artist := &ReleaseGroupArtist{}
    err := rows.Scan(&artist.Gid, &artist.Name)
    if err == nil {
      artists = append(artists, artist)
    }
  }

  return artists
}

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

  releaseGroup.Artists = FindArtistsByArtistCredit(artistCredit)

  return releaseGroup, nil
}

func ReleaseGroupHandler(w traffic.ResponseWriter, r *http.Request) {
  gid := r.URL.Query().Get("gid")
  releaseGroup, err := FindReleaseGroupByGid(gid)

  if err == nil {
    json.NewEncoder(w).Encode(releaseGroup)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    w.WriteHeader(http.StatusInternalServerError)
  }
}
