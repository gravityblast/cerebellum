package main

import (
  "fmt"
  "regexp"
)

var UUIDRegexp = regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")

type InvalidUUID struct {
  UUID string
}

func (err InvalidUUID) Error() string {
  return fmt.Sprintf("Invalid UUID: %v", err.UUID)
}

func isValidUUID(uuid string) bool {
  return UUIDRegexp.MatchString(uuid)
}

