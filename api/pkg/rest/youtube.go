package rest

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

var GoogleConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_ID"),
	ClientSecret: os.Getenv("GOOGLE_SECRET"),
	Endpoint:     google.Endpoint,
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URI"),
	Scopes:       []string{youtube.YoutubepartnerScope},
}

func (api *ConcertifyAPI) YoutubeLogin(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(20 * time.Minute)
	state := make(map[string]interface{})
	randomizedState := make([]byte, 16)
	rand.Read(randomizedState)
	state["random"] = randomizedState
	state["playlist_id"] = r.FormValue("id")
	json, err := json.Marshal(state)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encoded_json := base64.URLEncoding.EncodeToString(json)
	url := GoogleConfig.AuthCodeURL(encoded_json)
	cookie := http.Cookie{Name: "oauthstate", Value: encoded_json, Expires: expiration}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (api *ConcertifyAPI) YoutubeCallback(w http.ResponseWriter, r *http.Request) {
	state, err := r.Cookie("oauthstate")
	if err != nil || r.FormValue("state") != state.Value {
		http.Error(w, "Bad OAuth State", http.StatusInternalServerError)
		return
	}
	data, err := base64.URLEncoding.DecodeString(state.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var state_data map[string]interface{}
	err = json.Unmarshal(data, &state_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	playlist_id := state_data["playlist_id"].(string)
	token, err := GoogleConfig.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		http.Error(w, "Failed to Retrieve Token"+err.Error(), http.StatusInternalServerError)
		return
	}
	session_id, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "No Session Cookie", http.StatusInternalServerError)
		return
	}
	err = api.Session_Manager.SetYoutubeSession(session_id.Value, token)
	if err != nil {
		http.Error(w, "Failed to Set Session", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, os.Getenv("FRONTEND_URL")+"/confirm?playlist="+playlist_id, http.StatusPermanentRedirect)
}
