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
  SortName  string  `json:"sortName"`
  Comment   string  `json:"comment"`
  BeginDate *string `json:"beginDate"`
  EndDate   *string `json:"endDate"`
  Type      string  `json:"type"`
}

const FindArtistByGidQuery = `
  SELECT
    A.gid, A.name, A.sort_name, A.comment,
    A.begin_date_year, A.begin_date_month, A.begin_date_day,
    A.end_date_year, A.end_date_month, A.end_date_day, AT.name
  FROM
    artist A
  LEFT JOIN artist_type AT
    ON A.type = AT.id
  WHERE
    A.gid = $1 limit 1;`

func FindArtistByGid(gid string) (*Artist, error) {
  artist := &Artist{}

  if !isValidUUID(gid) {
    return artist, InvalidUUID{ gid }
  }

  var _type    *sql.NullString

  var beginDateYear   *sql.NullInt64
  var beginDateMonth  *sql.NullInt64
  var beginDateDay    *sql.NullInt64
  var endDateYear     *sql.NullInt64
  var endDateMonth    *sql.NullInt64
  var endDateDay      *sql.NullInt64

  row := DB.QueryRow(FindArtistByGidQuery, gid)
  err := row.Scan(&artist.Gid, &artist.Name, &artist.SortName, &artist.Comment,
                  &beginDateYear, &beginDateMonth, &beginDateDay,
                  &endDateYear, &endDateMonth, &endDateDay, &_type)

  if err != nil {
    return artist, err
  }

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

  if _type != nil {
    artist.Type = _type.String
  }

  return artist, nil
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
