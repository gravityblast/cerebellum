package main

import (
  "net/http"
  "fmt"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query()
  fmt.Fprint(w, params.Get("gid"))
}
