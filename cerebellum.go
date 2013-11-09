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
  // Artist:
  //   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56
  router.Get("/artists/:gid", ArtistHandler)

  // Release Group:
  //   /release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc
  //   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc
  router.Get("(/artists/:artist_gid)?/release-groups/:gid", ReleaseGroupHandler).
    AddBeforeFilter(CheckArtistFilter)

  // Release:
  //   /releases/79215cdf-4764-4dee-b0b9-fec1643df7c5
  //   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5
  router.Get("(/artists/:artist_gid)?/releases/:gid", ReleaseHandler)

  // Recordings:
  //   /releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings
  //   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings
  router.Get("(/artists/:artist_gid)?/releases/:release_gid/recordings", RecordingsHandler).
    AddBeforeFilter(CheckArtistFilter).
    AddBeforeFilter(CheckReleaseFilter)

  // Release Groups:
  //   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups
  router.Get("/artists/:artist_gid/release-groups", ReleaseGroupsHandler).
    AddBeforeFilter(CheckArtistFilter)

  // Recordings:
  //   /recordings/0c871a4a-efdf-47f8-98c2-cc277f806d2f
  //   /releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings/0c871a4a-efdf-47f8-98c2-cc277f806d2f
  //   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings/833f00e1-781f-4edd-90e4-e52712618862
  router.Get("((/artists/:artist_gid)?/releases/:release_gid)?/recordings/:gid", RecordingHandler).
    AddBeforeFilter(CheckArtistFilter).
    AddBeforeFilter(CheckReleaseFilter)
}

func main() {
  router.Run()
}
