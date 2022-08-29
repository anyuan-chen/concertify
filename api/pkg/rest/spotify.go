package rest

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"time"

	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

var scopes = spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserTopRead, spotifyauth.ScopeUserLibraryModify, spotifyauth.ScopePlaylistModifyPublic, spotifyauth.ScopePlaylistModifyPrivate, spotifyauth.ScopeUserLibraryRead)
var auth = spotifyauth.New(spotifyauth.WithRedirectURL(os.Getenv("SPOTIFY_REDIRECT_URI")), scopes)

func (api *ConcertifyAPI) SpotifyLogin(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(20 * time.Minute)
	state := make(map[string]interface{})
	randomizedState := make([]byte, 16)
	rand.Read(randomizedState)
	state["random"] = randomizedState
	json, err := json.Marshal(state)
	if err != nil {
		http.Error(w, "A JSON Encoding error has been encountered.", http.StatusInternalServerError)
		return
	}
	encoded_json := base64.URLEncoding.EncodeToString(json)
	url := auth.AuthURL(encoded_json)
	cookie := http.Cookie{Name: "oauthstate", Value: encoded_json, Expires: expiration}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (api *ConcertifyAPI) SpotifyCallback(w http.ResponseWriter, r *http.Request) {
	state, err := r.Cookie("oauthstate")
	if err != nil || r.FormValue("state") != state.Value {
		http.Error(w, "Bad OAuth State", http.StatusInternalServerError)
	}
	token, err := auth.Token(r.Context(), state.Value, r)
	if err != nil {
		http.Error(w, "Failed to Retrieve Token", http.StatusInternalServerError)
	}
	id, err := api.Session_Manager.SetSpotifySession(token)
	if err != nil {
		http.Error(w, "Problem with the session management service: "+err.Error(), http.StatusInternalServerError)
	}
	cookie := http.Cookie{Name: "session_id", Path: "/", Value: id, Secure: true, Expires: time.Now().Add(time.Hour * 24 * 7)}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, os.Getenv("FRONTEND_URL")+"/select", http.StatusPermanentRedirect)
}
