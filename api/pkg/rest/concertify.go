package rest

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/zmb3/spotify/v2"
	"google.golang.org/api/youtube/v3"
)

func (api *ConcertifyAPI) GetAllPlaylists(w http.ResponseWriter, r *http.Request) {
	session_id_cookie, err := r.Cookie("session_id")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"bad cookie", http.StatusUnauthorized)
		return
	}
	session_id := session_id_cookie.Value
	spotify_token, err := api.Session_Manager.GetSpotifySession(session_id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"session manager can't find", http.StatusUnauthorized)
		return
	}
	pageNumber, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"bad page number", http.StatusBadRequest)
		return
	}
	playlists, err := api.ConcertifyCore.GetAllPlaylists(context.Background(), spotify_token, pageNumber)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"failed to retrieve", http.StatusInternalServerError)
		return
	}
	playlists_json, err := json.Marshal(playlists)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(playlists_json)
}

func (api *ConcertifyAPI) ViewPlaylist(w http.ResponseWriter, r *http.Request) {
	session_id_cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	session_id := session_id_cookie.Value
	spotify_token, err := api.Session_Manager.GetSpotifySession(session_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	playlist_id := spotify.ID(r.FormValue("playlist_id"))
	playlist, err := api.ConcertifyCore.GetPlaylist(context.Background(), spotify_token, playlist_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responses := make([][]*youtube.SearchResult, 0, len(playlist.Tracks.Tracks))
	for _, track := range playlist.Tracks.Tracks {
		trackName := track.Track.Name
		searchListResponse, err := api.ConcertifyCore.GetYoutubeVideoFromSpotify(trackName, 3)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responses = append(responses, searchListResponse.Items)
	}

	type ConcertifyPlaylist struct {
		SearchResponse [][]*youtube.SearchResult `json:"youtube_search_response"`
		Playlist       []spotify.PlaylistTrack   `json:"spotify_playlist"`
	}

	combined := &ConcertifyPlaylist{
		SearchResponse: responses,
		Playlist:       playlist.Tracks.Tracks,
	}

	combined_json, err := json.Marshal(combined)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(combined_json)
}
