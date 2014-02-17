package main

import (
	"database/sql"
	"encoding/json"
	"github.com/pilu/cerebellum/models"
	"github.com/pilu/cerebellum/models/recording"
	"github.com/pilu/traffic"
	"net/http"
)

func RecordingHandler(w traffic.ResponseWriter, r *traffic.Request) {
	releaseId := r.URL.Query().Get("release_id")
	id := r.URL.Query().Get("id")

	var Recording *models.Recording
	var err error

	if releaseId != "" {
		Recording, err = recording.ByReleaseIdAndId(releaseId, id)
	} else {
		Recording, err = recording.ById(id)
	}

	if err == nil {
		json.NewEncoder(w).Encode(Recording)
	} else if err == sql.ErrNoRows {
		RecordingNotFoundHandler(w, r)
	} else if _, ok := err.(models.InvalidUUID); ok {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		panic(err)
	}
}
