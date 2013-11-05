package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/pilu/traffic"
)

type Artist struct {
  Gid       string  `json:"gid"`
  Name      string  `json:"name"`
  Comment   string  `json:"comment"`
  BeginDate *string `json:"beginDate"`
  EndDate   *string `json:"endDate"`
  SortName  string  `json:"sortName"`
  Type      string  `json:"type"`
}

const FindArtistByGidQuery = `
  SELECT
    A.gid, AN.name, A.comment,
    A.begin_date_year, A.begin_date_month, A.begin_date_day,
    A.end_date_year, A.end_date_month, A.end_date_day,
    ASN.name, AT.name
  FROM
    artist A
  LEFT JOIN artist_name AN
    ON A.name = AN.id
  LEFT JOIN artist_name ASN
    ON A.sort_name = ASN.id
  LEFT JOIN artist_type AT
    ON A.type = AT.id
  WHERE
    A.gid = $1 limit 1;`

func FindArtistByGid(gid string) (*Artist, error) {
  artist := &Artist{}

  if !isValidUUID(gid) {
    return artist, InvalidUUID{ gid }
  }

  var name     *sql.NullString
  var sortName *sql.NullString
  var _type    *sql.NullString

  var beginDateYear   *sql.NullInt64
  var beginDateMonth  *sql.NullInt64
  var beginDateDay    *sql.NullInt64
  var endDateYear     *sql.NullInt64
  var endDateMonth    *sql.NullInt64
  var endDateDay      *sql.NullInt64

  err := DB.QueryRow(FindArtistByGidQuery, gid).Scan(&artist.Gid, &name, &artist.Comment,
                                                      &beginDateYear, &beginDateMonth, &beginDateDay,
                                                      &endDateYear, &endDateMonth, &endDateDay,
                                                      &sortName, &_type)

  if beginDateYear != nil {
    date := fmt.Sprintf("%d", beginDateYear.Int64)

    if beginDateMonth != nil  {
      date = fmt.Sprintf("%s-%02d", date, beginDateMonth.Int64)
    }

    if beginDateDay != nil  {
      date = fmt.Sprintf("%s-%02d", date, beginDateDay.Int64)
    }

    artist.BeginDate = &date
  }

  if endDateYear != nil {
    date := fmt.Sprintf("%d", endDateYear.Int64)

    if endDateMonth != nil  {
      date = fmt.Sprintf("%s-%02d", date, endDateMonth.Int64)
    }

    if endDateDay != nil  {
      date = fmt.Sprintf("%s-%02d", date, endDateDay.Int64)
    }

    artist.EndDate = &date
  }

  if name != nil {
    artist.Name = name.String
  }

  if sortName != nil {
    artist.SortName = sortName.String
  }

  if _type != nil {
    artist.Type = _type.String
  }

  return artist, err
}

func ArtistHandler(w traffic.ResponseWriter, r *http.Request) {
  gid := r.URL.Query().Get("gid")
  artist, err := FindArtistByGid(gid)

  if err == nil {
    json.NewEncoder(w).Encode(artist)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    w.WriteHeader(http.StatusInternalServerError)
  }
}
