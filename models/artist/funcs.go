package artist

import (
  "fmt"
  "database/sql"
  "github.com/pilu/cerebellum/models"
)

func ByGid(gid string) (*models.Artist, error) {
  artist := &models.Artist{}

  if !models.IsValidUUID(gid) {
    return artist, models.InvalidUUID{ gid }
  }

  var _type    *sql.NullString

  var beginDateYear   *sql.NullInt64
  var beginDateMonth  *sql.NullInt64
  var beginDateDay    *sql.NullInt64
  var endDateYear     *sql.NullInt64
  var endDateMonth    *sql.NullInt64
  var endDateDay      *sql.NullInt64

  row := models.DB.QueryRow(queryByGid, gid)
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

    artist.BeginDate = date
  }

  if endDateYear != nil {
    date := fmt.Sprintf("%d", endDateYear.Int64)

    if endDateMonth != nil  {
      date = fmt.Sprintf("%s-%02d", date, endDateMonth.Int64)
    }

    if endDateDay != nil  {
      date = fmt.Sprintf("%s-%02d", date, endDateDay.Int64)
    }

    artist.EndDate = date
  }

  if _type != nil {
    artist.Type = _type.String
  }

  return artist, nil
}

func AllByArtistCredit(artistCredit int) []*models.ReleaseArtist {
  artists := make([]*models.ReleaseArtist, 0)

  rows, err := models.DB.Query(queryAllByArtistCredit, artistCredit)
  if err != nil {
    return artists
  }

  for rows.Next() {
    artist := &models.ReleaseArtist{}
    err := rows.Scan(&artist.Gid, &artist.Name)
    if err == nil {
      artists = append(artists, artist)
    }
  }

  return artists
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
