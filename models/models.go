package models

import (
  "database/sql"
)

var DB *sql.DB

type Artist struct {
  Gid       string `json:"gid"`
  Name      string `json:"name"`
  SortName  string `json:"sortName"`
  Comment   string `json:"comment"`
  BeginDate string `json:"beginDate"`
  EndDate   string `json:"endDate"`
  Type      string `json:"type"`
}

type ReleaseArtist struct {
  Gid   string `json:"gid"`
  Name  string `json:"name"`
}

type ReleaseGroup struct {
  Gid               string           `json:"gid"`
  Name              string           `json:"name"`
  Comment           string           `json:"comment"`
  FirstReleaseDate  string           `json:"firstReleaseDate"`
  Type              string           `json:"type"`
  ArtistCredit      int              `json:"-"`
  Artists           []*ReleaseArtist `json:"artists,omitempty"`
}

type Release struct {
  Gid               string           `json:"gid"`
  Name              string           `json:"name"`
  Comment           string           `json:"comment"`
  Status            string           `json:"status"`
  Packaging         string           `json:"packaging"`
  Artists           []*ReleaseArtist `json:"artists"`
}


type Recording struct {
  Gid               string           `json:"gid"`
  Name              string           `json:"name"`
  Comment           string           `json:"comment"`
  Length            int              `json:"length"`
  Artists           []*ReleaseArtist `json:"artists"`
}
