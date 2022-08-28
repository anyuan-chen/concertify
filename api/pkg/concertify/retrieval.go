package concertify

import (
	"context"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

func (c *ConcertifyCore) GetPopularityScore(ctx context.Context, token *oauth2.Token) error {
	client := spotify.New(c.Authenticator.Client(context.Background(), token))
	tracks, err := client.CurrentUsersTopTracks(context.Background(), spotify.Limit(50))
	if err != nil {
		return err
	}
	var average float64 = 0
	for _, track := range tracks.Tracks {
		average += float64(track.Popularity)
	}
	average = average / (float64(len(tracks.Tracks)))
	return nil
}
