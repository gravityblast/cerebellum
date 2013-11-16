package models

import (
  "fmt"
  "regexp"
  "database/sql"
)

var UUIDRegexp = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

type InvalidUUID struct {
  UUID string
}

func (err InvalidUUID) Error() string {
  return fmt.Sprintf("Invalid UUID: %v", err.UUID)
}

func IsValidUUID(uuid string) bool {
  return UUIDRegexp.MatchString(uuid)
}

func DatesToString(year, month, day *sql.NullInt64) string {
  var date string

  if year != nil {
    date = fmt.Sprintf("%d", year.Int64)

    if month != nil {
      date = fmt.Sprintf("%s-%02d", date, month.Int64)
    }

    if day != nil {
      date = fmt.Sprintf("%s-%02d", date, day.Int64)
    }
  }

  return date
}
