package main

import (
  "net/http"
  "encoding/json"
  "database/sql"
)

type Artist struct {
  Gid string `json:"gid"`
  Name string `json:"name"`
  SortName string `json:"sortName"`
}

const FindArtistByGidQuery = `select A.gid as gid, AN.name as artist_name, ASN.name as sort_name from
  artist A, artist_name AN, artist_name ASN
  where A.gid = $1 AND A.name = AN.id AND A.sort_name = ASN.id limit 1;`

func FindArtistByGid(gid string) (*Artist, error) {
  artist := &Artist{}

  if !isValidUUID(gid) {
    return artist, InvalidUUID{ gid }
  }
  err := DB.QueryRow(FindArtistByGidQuery, gid).Scan(&artist.Gid, &artist.Name, &artist.SortName)

  return artist, err
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
  gid := r.URL.Query().Get("gid")
  artist, err := FindArtistByGid(gid)

  if err == nil {
    json.NewEncoder(w).Encode(artist)
  } else if err == sql.ErrNoRows {
    w.WriteHeader(http.StatusNotFound)
  } else if _, ok := err.(InvalidUUID); ok {
    w.WriteHeader(http.StatusBadRequest)
  } else {
    w.WriteHeader(http.StatusInternalServerError)
  }
}
