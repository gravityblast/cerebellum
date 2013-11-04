package main

import (
  "os"
  "fmt"
  "regexp"
  "net/http"
  "encoding/json"
  "database/sql"
  _ "github.com/bmizerany/pq"
  "github.com/pilu/traffic"
)

const VERSION = "0.1.0"
var router *traffic.Router

var UUIDRegexp *regexp.Regexp
var DB *sql.DB

type InvalidUUID struct {
  UUID string
}

func (err InvalidUUID) Error() string {
  return fmt.Sprintf("Invalid UUID: %v", err.UUID)
}

func init() {
  UUIDRegexp = regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")

  var err error
  DB, err = sql.Open("postgres", os.Getenv("DB"))
  if err != nil {
    panic(err)
  }

  router = traffic.New()
  router.NotFoundHandler = NotFoundHandler
  router.AddBeforeFilter(SetDefaultHeaders)

  router.Get("/", RootHandler)
  router.Get("/artists/:gid", ArtistHandler)
}

func NotFoundHandler(w traffic.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(map[string]string{
    "error": "Not Found",
  })
}

func isValidUUID(uuid string) bool {
  return UUIDRegexp.MatchString(uuid)
}

func SetDefaultHeaders(w traffic.ResponseWriter, r *http.Request) bool {
  w.Header().Set("Cerebellum-Version", VERSION)
  w.Header().Set("Content-Type", "application/json; charset=utf-8")

  return true
}

func main() {
  router.Run()
}
