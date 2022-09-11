package rest

import (
	"context"
	"encoding/json"
	"io"
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
		log.Println(err)
		http.Error(w, err.Error()+"no session id", http.StatusUnauthorized)
		return
	}
	session_id := session_id_cookie.Value
	spotify_token, err := api.Session_Manager.GetSpotifySession(session_id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"no spotify token", http.StatusUnauthorized)
		return
	}
	playlist_id := spotify.ID(r.FormValue("playlist"))
	playlist, err := api.ConcertifyCore.GetPlaylist(context.Background(), spotify_token, playlist_id)
	if err != nil {
		log.Println(err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"no youtube token", http.StatusUnauthorized)
		return
	}
	responses := make([][]*youtube.SearchResult, 0, len(playlist.Tracks.Tracks))
	for _, track := range playlist.Tracks.Tracks {
		trackName := track.Track.Name
		artistName := track.Track.Artists[0].Name
		searchListResponse, err := api.ConcertifyCore.GetYoutubeVideoFromSpotify(trackName+" "+artistName, 3)
		if err != nil {
			log.Println(err)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responses = append(responses, searchListResponse.Items)
	}

	// type ConcertifyPlaylist struct {
	// 	SearchResponse [][]*youtube.SearchResult `json:"youtube_search_response"`
	// 	Playlist       []spotify.PlaylistTrack   `json:"spotify_playlist"`
	// }

	combined := make([]ConcertifyPlaylistItem, 0, len(playlist.Tracks.Tracks))
	for i, track := range playlist.Tracks.Tracks {
		combined = append(combined, ConcertifyPlaylistItem{
			FullTrack:             track.Track,
			YoutubeSearchResponse: responses[i],
		})
	}
	combined_json, err := json.Marshal(combined)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(combined_json)
}

func (api *ConcertifyAPI) MakePlaylist(w http.ResponseWriter, r *http.Request) {
	session_id_cookie, err := r.Cookie("session_id")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"no session id", http.StatusUnauthorized)
		return
	}
	session_id := session_id_cookie.Value
	youtube_token, err := api.Session_Manager.GetYoutubeSession(session_id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"no youtube token", http.StatusUnauthorized)
		return
	}
	id, err := api.ConcertifyCore.CreateYoutubePlaylist(youtube_token)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"creating playlist", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	items, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"reading the body", http.StatusInternalServerError)
		return
	}
	var playlistInformation []ConcertifyPlaylistItem
	err = json.Unmarshal(items, &playlistInformation)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"parsing the body into json", http.StatusInternalServerError)
		return
	}
	playlistIds := make([]string, 0, len(playlistInformation))
	for _, item := range playlistInformation {
		playlistIds = append(playlistIds, item.YoutubeSearchResponse[0].Id.VideoId)
	}
	err = api.ConcertifyCore.InsertItemsIntoPlaylist(youtube_token, playlistIds, id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"inserting the items into the playlist", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	type PlaylistId struct {
		Id string `json:"id"`
	}
	playlistId := PlaylistId{
		Id: id,
	}
	playlistId_json, err := json.Marshal(playlistId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"marshalling the playlist id", http.StatusInternalServerError)
		return
	}
	w.Write(playlistId_json)
}

func (api *ConcertifyAPI) GetYoutubeVideo(w http.ResponseWriter, r *http.Request) {
	session_id_cookie, err := r.Cookie("session_id")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"no session id", http.StatusUnauthorized)
		return
	}
	session_id := session_id_cookie.Value
	youtube_token, err := api.Session_Manager.GetYoutubeSession(session_id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"no youtube token", http.StatusUnauthorized)
		return
	}
	searchListResponse, err := api.ConcertifyCore.GetVideoFromLink(context.Background(), youtube_token, r.FormValue("query"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"getting youtube video", http.StatusInternalServerError)
		return
	}
	searchListResponse_json, err := json.Marshal(searchListResponse)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error()+"marshalling the youtube video", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(searchListResponse_json)
}
