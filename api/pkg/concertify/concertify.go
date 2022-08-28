package concertify

import spotifyauth "github.com/zmb3/spotify/v2/auth"

type ConcertifyCore struct {
	Authenticator *spotifyauth.Authenticator
}

func NewConcertifyCore(Authenticator *spotifyauth.Authenticator) ConcertifyCore {
	return ConcertifyCore{
		Authenticator: Authenticator,
	}
}
