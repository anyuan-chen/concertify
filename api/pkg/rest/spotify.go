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
var auth = spotifyauth.New(spotifyauth.WithRedirectURL(os.Getenv("REDIRECT_URL")), scopes)

func (api *ConcertifyAPI) SpotifyLogin(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(20 * time.Minute)
	state := make(map[string]interface{})
	randomizedState := make([]byte, 16)
	rand.Read(randomizedState)
	state["random"] = randomizedState
	json, err := json.Marshal(state)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encoded_json := base64.URLEncoding.EncodeToString(json)
	url := auth.AuthURL(encoded_json)
	cookie := http.Cookie{Name: "oauth_state", Value: encoded_json, Expires: expiration}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
