package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zmb3/spotify/v2"
)

func (api *ConcertifyAPI) GetAllPlaylists(w http.ResponseWriter, r *http.Request) {
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
	playlists, err := api.ConcertifyCore.GetAllPlaylists(context.Background(), spotify_token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	playlists_json, err := json.Marshal(playlists)
	if err != nil {
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
	playlist_json, err := json.Marshal(playlist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(playlist_json)
}
