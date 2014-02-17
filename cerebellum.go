package main

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"github.com/pilu/cerebellum/models"
	"github.com/pilu/traffic"
)

const VERSION = "0.1.0"

var (
	router *traffic.Router
	DB     *sql.DB
)

func initDatabase() {
	var err error

	dbUser := traffic.GetVar("database.user")
	dbName := traffic.GetVar("database.name")
	dbHost := traffic.GetVar("database.host")
	dbPass := traffic.GetVar("database.pass")
	dbSSLMode := traffic.GetVar("database.sslmode")

	dbString := fmt.Sprintf("user=%s dbname=%s host=%s password=%s sslmode=%s", dbUser, dbName, dbHost, dbPass, dbSSLMode)

	DB, err = sql.Open("postgres", dbString)
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
	router.ErrorHandler = ErrorHandler
	router.AddBeforeFilter(SetDefaultHeaders)

	router.Get("/", RootHandler)
	// Artist:
	//   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56
	router.Get("/artists/:id/?", ArtistHandler)

	// Release Group:
	//   /release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc
	//   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc
	router.Get("(/artists/:artist_id)?/release-groups/:id/?", ReleaseGroupHandler).
		AddBeforeFilter(CheckArtistFilter)

	// Release Groups:
	//   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups
	router.Get("/artists/:artist_id/release-groups/?", ReleaseGroupsHandler).
		AddBeforeFilter(CheckArtistFilter)

	// Release:
	//   /releases/79215cdf-4764-4dee-b0b9-fec1643df7c5
	//   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5
	router.Get("(/artists/:artist_id)?/releases/:id/?", ReleaseHandler).
		AddBeforeFilter(CheckArtistFilter)

	// Releases:
	//   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases
	router.Get("/artists/:artist_id/releases/?", ReleasesHandler).
		AddBeforeFilter(CheckArtistFilter)

	// Releases:
	//   /release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc/releases
	//   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc/releases
	router.Get("(/artists/:artist_id)?/release-groups/:release_group_id/releases/?", ReleasesHandler).
		AddBeforeFilter(CheckArtistFilter).
		AddBeforeFilter(CheckReleaseGroupFilter)

	// Recordings:
	//   /releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings
	//   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings
	router.Get("(/artists/:artist_id)?/releases/:release_id/recordings/?", RecordingsHandler).
		AddBeforeFilter(CheckArtistFilter).
		AddBeforeFilter(CheckReleaseFilter)

	// Recording:
	//   /recordings/0c871a4a-efdf-47f8-98c2-cc277f806d2f
	//   /releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings/0c871a4a-efdf-47f8-98c2-cc277f806d2f
	//   /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings/833f00e1-781f-4edd-90e4-e52712618862
	router.Get("((/artists/:artist_id)?/releases/:release_id)?/recordings/:id/?", RecordingHandler).
		AddBeforeFilter(CheckArtistFilter).
		AddBeforeFilter(CheckReleaseFilter)
}

func main() {
	router.Run()
}
