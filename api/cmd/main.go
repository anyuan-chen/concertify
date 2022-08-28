package main

import (
	"net/http"
	"os"

	"github.com/anyuan-chen/concertify/api/pkg/concertify"
	"github.com/anyuan-chen/concertify/api/pkg/rest"
	"github.com/anyuan-chen/concertify/api/pkg/session_manager"
	"github.com/gorilla/mux"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type Server struct {
	r *mux.Router
}

func main() {
	auth := spotifyauth.New(spotifyauth.WithRedirectURL(os.Getenv("REDIRECT_URL")), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
	concertify_core := concertify.NewConcertifyCore(auth)

	session_manager := session_manager.CreateManager()
	concertify_api := rest.NewConcertifyAPI(session_manager, &concertify_core)

	r := mux.NewRouter()
	r.HandleFunc("/spotify/login", concertify_api.SpotifyLogin)
	r.HandleFunc("/spotify/callback", concertify_api.SpotifyCallback)
	r.HandleFunc("/youtube/login", concertify_api.YoutubeLogin)
	r.HandleFunc("/youtube/callback", concertify_api.YoutubeCallback)
	http.Handle("/", &Server{r: r})
	http.ListenAndServe(":8080", nil)
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	if req.Method == "OPTIONS" {
		return
	}
	s.r.ServeHTTP(rw, req)
}
