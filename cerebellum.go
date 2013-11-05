package main

import (
  "os"
  "net/http"
  "encoding/json"
  "database/sql"
  _ "github.com/bmizerany/pq"
  "github.com/pilu/traffic"
)

const VERSION = "0.1.0"

var (
  router *traffic.Router
  DB     *sql.DB
)

func NotFoundHandler(w traffic.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(map[string]string{
    "error": "Not Found",
  })
}

func SetDefaultHeaders(w traffic.ResponseWriter, r *http.Request) bool {
  w.Header().Set("Cerebellum-Version", VERSION)
  w.Header().Set("Content-Type", "application/json; charset=utf-8")

  return true
}

func init() {
  var err error
  DB, err = sql.Open("postgres", os.Getenv("DB"))
  if err != nil {
    panic(err)
  }

  router = traffic.New()
  router.NotFoundHandler = NotFoundHandler
  router.AddBeforeFilter(SetDefaultHeaders)

  router.Get("/", RootHandler)
  router.Get("/artists/:gid",         ArtistHandler)
  router.Get("/release-groups/:gid",  ReleaseGroupHandler)
}

func main() {
  router.Run()
}
