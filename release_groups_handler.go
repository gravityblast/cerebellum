package main

import (
	"database/sql"
	"encoding/json"
	"github.com/pilu/cerebellum/models"
	"github.com/pilu/cerebellum/models/releasegroup"
	"github.com/pilu/traffic"
	"net/http"
)

func ReleaseGroupsHandler(w traffic.ResponseWriter, r *traffic.Request) {
	artistId := r.URL.Query().Get("artist_id")

	ReleaseGroups, err := releasegroup.AllByArtistId(artistId)

	if err == nil {
		json.NewEncoder(w).Encode(ReleaseGroups)
	} else if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
	} else if _, ok := err.(models.InvalidUUID); ok {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		panic(err)
	}
}
