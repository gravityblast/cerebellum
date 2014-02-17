package models

import (
  "database/sql"
)

var DB *sql.DB

type Artist struct {
  Id       string `json:"id"`
  Name      string `json:"name"`
  SortName  string `json:"sortName"`
  Comment   string `json:"comment"`
  BeginDate string `json:"beginDate"`
  EndDate   string `json:"endDate"`
  Type      string `json:"type"`
}

type ReleaseArtist struct {
  Id   string `json:"id"`
  Name  string `json:"name"`
}

type ReleaseGroup struct {
  Id               string           `json:"id"`
  Name              string           `json:"name"`
  Comment           string           `json:"comment"`
  FirstReleaseDate  string           `json:"firstReleaseDate"`
  Type              string           `json:"type"`
  ArtistCredit      int              `json:"-"`
  Artists           []*ReleaseArtist `json:"artists,omitempty"`
}

type Release struct {
  Id               string           `json:"id"`
  Name              string           `json:"name"`
  Comment           string           `json:"comment"`
  Date              string           `json:"date,omitempty"`
  Status            string           `json:"status"`
  Type              string           `json:"type,omitempty"`
  Packaging         string           `json:"packaging"`
  Artists           []*ReleaseArtist `json:"artists,omitempty"`
}


type Recording struct {
  Id               string           `json:"id"`
  Name              string           `json:"name"`
  Comment           string           `json:"comment"`
  Length            int              `json:"length"`
  Artists           []*ReleaseArtist `json:"artists,omitempty"`
}

