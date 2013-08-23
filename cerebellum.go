package main

import (
  "net/http"
  "github.com/pilu/traffic"
  "fmt"
  "os"
  "regexp"
)

const VERSION = "0.1.0"
var router *traffic.Router

var UUIDRegexp *regexp.Regexp

func init() {
  UUIDRegexp = regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")

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
