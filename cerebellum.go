package main

import (
  "net/http"
  "github.com/pilu/traffic"
  "fmt"
  "os"
  "regexp"
  "database/sql"
  _ "github.com/bmizerany/pq"
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
  router.AddBeforeFilter(SetDefaultHeaders)

  router.Get("/", RootHandler)
  router.Get("/artists/:gid", ArtistHandler)
}

func isValidUUID(uuid string) bool {
  return UUIDRegexp.MatchString(uuid)
}

func SetDefaultHeaders(w http.ResponseWriter, r *http.Request) bool {
  w.Header().Add("Cerebellum-Version", VERSION)
  w.Header().Add("Content-Type", "application/json")

  return true
}

func main() {
  http.Handle("/", router)

  port := os.Getenv("PORT")
  host := os.Getenv("HOST")

  err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
  if err != nil {
    panic(err)
  }
}
