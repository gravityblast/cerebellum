package main

import (
  "os"
  "database/sql"
  _ "github.com/bmizerany/pq"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
)

const VERSION = "0.1.0"

var (
  router *traffic.Router
  DB     *sql.DB
)

func initDatabase() {
  var err error
  DB, err = sql.Open("postgres", os.Getenv("DB"))
  if err != nil {
    panic(err)
  }

  err = DB.Ping()
  if err != nil {
    panic(err)
  }

  models.DB = DB
}

func init() {
  initDatabase()

  router = traffic.New()
  router.NotFoundHandler = NotFoundHandler
  router.ErrorHandler    = ErrorHandler
  router.AddBeforeFilter(SetDefaultHeaders)

  router.Get("/", RootHandler)
  router.Get("/artists/:gid", ArtistHandler)
  router.Get("(/artists/:artist_gid)?/release-groups/:gid", ReleaseGroupHandler)
  router.Get("(/artists/:artist_gid)?/releases/:gid", ReleaseHandler)
  router.Get("(/artists/:artist_gid)?/releases/:release_gid/recordings", RecordingsHandler).
    AddBeforeFilter(CheckArtistFilter).
    AddBeforeFilter(CheckReleaseFilter)
  router.Get("/artists/:artist_gid/release-groups", ReleaseGroupsHandler)
  router.Get("/recordings/:gid", RecordingHandler)
}

func main() {
  router.Run()
}
