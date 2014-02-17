package main

import (
	"database/sql"
	"encoding/json"
	"github.com/pilu/cerebellum/models"
	"github.com/pilu/cerebellum/models/release"
	"github.com/pilu/traffic"
	"net/http"
)

func ReleaseHandler(w traffic.ResponseWriter, r *traffic.Request) {
	artistId := r.URL.Query().Get("artist_id")
	id := r.URL.Query().Get("id")

	var Release *models.Release
	var err error

	if artistId == "" {
		Release, err = release.ById(id)
	} else {
		Release, err = release.ByArtistIdAndId(artistId, id)
	}

	if err == nil {
		json.NewEncoder(w).Encode(Release)
	} else if err == sql.ErrNoRows {
		ReleaseNotFoundHandler(w, r)
	} else if _, ok := err.(models.InvalidUUID); ok {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		panic(err)
	}
}
