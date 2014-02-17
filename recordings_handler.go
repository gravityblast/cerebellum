package main

import (
	"database/sql"
	"encoding/json"
	"github.com/pilu/cerebellum/models"
	"github.com/pilu/cerebellum/models/recording"
	"github.com/pilu/traffic"
	"net/http"
)

func RecordingsHandler(w traffic.ResponseWriter, r *traffic.Request) {
	releaseId := r.URL.Query().Get("release_id")

	var recordings []*models.Recording
	var err error

	recordings, err = recording.AllByReleaseId(releaseId)

	if err == nil {
		json.NewEncoder(w).Encode(recordings)
	} else if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
	} else if _, ok := err.(models.InvalidUUID); ok {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		panic(err)
	}
}
