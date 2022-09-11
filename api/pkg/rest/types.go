package rest

import (
	"github.com/zmb3/spotify/v2"
	"google.golang.org/api/youtube/v3"
)

type ConcertifyPlaylistItem struct {
	spotify.FullTrack
	YoutubeSearchResponse []*youtube.SearchResult `json:"youtube_search_response"`
}
