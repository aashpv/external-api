package info

import (
	"encoding/json"
	"external-api/internal/models"
	"net/http"
)

type SongInfoGetter interface {
	GetSongDetails(group, song string) (*models.SongDetails, error)
}

func New(getter SongInfoGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group := r.URL.Query().Get("group")
		song := r.URL.Query().Get("song")

		if group == "" || song == "" {
			http.Error(w, "Group and song parameters are required", http.StatusBadRequest)
			return
		}

		songDetails, err := getter.GetSongDetails(group, song)
		if err != nil {
			http.Error(w, "Song not found", http.StatusNotFound)
			return
		}

		response := models.SongDetails{
			ReleaseDate: songDetails.ReleaseDate,
			Text:        songDetails.Text,
			Link:        songDetails.Link,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
