package main

import (
	"encoding/json"
	"github.com/pilu/cerebellum/models"
	"github.com/pilu/cerebellum/models/artist"
	"github.com/pilu/cerebellum/models/release"
	"github.com/pilu/cerebellum/models/releasegroup"
	"github.com/pilu/traffic"
	"net/http"
)

func ErrorHandler(w traffic.ResponseWriter, r *traffic.Request, err interface{}) {
	json.NewEncoder(w).Encode(map[string]string{
		"error": "something went wrong",
	})
}

func NotFoundHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "not found",
	})
}

func SetDefaultHeaders(w traffic.ResponseWriter, r *traffic.Request) {
	w.Header().Set("Cerebellum-Version", VERSION)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func ArtistNotFoundHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "artist not found",
	})
}

func ReleaseGroupNotFoundHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "release group not found",
	})
}

func ReleaseNotFoundHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "release not found",
	})
}

func RecordingNotFoundHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "recording not found",
	})
}

func CheckArtistFilter(w traffic.ResponseWriter, r *traffic.Request) {
	artistId := r.URL.Query().Get("artist_id")
	if artistId == "" {
		return
	}

	if !models.IsValidUUID(artistId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !artist.Exists(artistId) {
		ArtistNotFoundHandler(w, r)
		return
	}
}

func CheckReleaseFilter(w traffic.ResponseWriter, r *traffic.Request) {
	artistId := r.URL.Query().Get("artist_id")
	releaseId := r.URL.Query().Get("release_id")

	if releaseId == "" {
		return
	}

	if !models.IsValidUUID(releaseId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if artistId != "" {
		if !models.IsValidUUID(artistId) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if artist.HasRelease(artistId, releaseId) {
			return
		}

		ReleaseNotFoundHandler(w, r)
		return
	}

	if !release.Exists(releaseId) {
		ReleaseNotFoundHandler(w, r)
		return
	}
}

func CheckReleaseGroupFilter(w traffic.ResponseWriter, r *traffic.Request) {
	artistId := r.URL.Query().Get("artist_id")
	releaseGroupId := r.URL.Query().Get("release_group_id")

	if releaseGroupId == "" {
		return
	}

	if !models.IsValidUUID(releaseGroupId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if artistId != "" {
		if !models.IsValidUUID(artistId) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if artist.HasReleaseGroup(artistId, releaseGroupId) {
			return
		}

		ReleaseGroupNotFoundHandler(w, r)
		return
	}

	if !releasegroup.Exists(releaseGroupId) {
		ReleaseGroupNotFoundHandler(w, r)
		return
	}
}
