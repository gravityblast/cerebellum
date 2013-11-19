package main

import (
  "net/http"
  "encoding/json"
  "github.com/pilu/traffic"
  "github.com/pilu/cerebellum/models"
  "github.com/pilu/cerebellum/models/artist"
  "github.com/pilu/cerebellum/models/release"
  "github.com/pilu/cerebellum/models/releasegroup"
)

func ErrorHandler(w traffic.ResponseWriter, r *traffic.Request, err interface{}) {
  json.NewEncoder(w).Encode(map[string]string{
    "error": "something went wrong",
  })
}

func NotFoundHandler(w traffic.ResponseWriter, r *traffic.Request) {
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
  artistGid := r.URL.Query().Get("artist_gid")
  if artistGid == "" {
    return
  }

  if !models.IsValidUUID(artistGid) {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  if !artist.Exists(artistGid) {
    ArtistNotFoundHandler(w, r)
    return
  }
}

func CheckReleaseFilter(w traffic.ResponseWriter, r *traffic.Request) {
  artistGid   := r.URL.Query().Get("artist_gid")
  releaseGid  := r.URL.Query().Get("release_gid")

  if releaseGid == "" {
    return
  }

  if !models.IsValidUUID(releaseGid) {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  if artistGid != "" {
    if !models.IsValidUUID(artistGid) {
      w.WriteHeader(http.StatusBadRequest)
      return
    }

    if artist.HasRelease(artistGid, releaseGid) {
      return
    }

    ReleaseNotFoundHandler(w, r)
    return
  }

  if !release.Exists(releaseGid) {
    ReleaseNotFoundHandler(w, r)
    return
  }
}

func CheckReleaseGroupFilter(w traffic.ResponseWriter, r *traffic.Request) {
  artistGid       := r.URL.Query().Get("artist_gid")
  releaseGroupGid := r.URL.Query().Get("release_group_gid")

  if releaseGroupGid == "" {
    return
  }

  if !models.IsValidUUID(releaseGroupGid) {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  if artistGid != "" {
    if !models.IsValidUUID(artistGid) {
      w.WriteHeader(http.StatusBadRequest)
      return
    }

    if artist.HasReleaseGroup(artistGid, releaseGroupGid) {
      return
    }

    ReleaseGroupNotFoundHandler(w, r)
    return
  }

  if !releasegroup.Exists(releaseGroupGid) {
    ReleaseGroupNotFoundHandler(w, r)
    return
  }
}
