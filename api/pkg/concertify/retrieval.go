package concertify

import (
	"context"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

func (c *ConcertifyCore) GetAllPlaylists(ctx context.Context, token *oauth2.Token) (*spotify.SimplePlaylistPage, error) {
	client := spotify.New(c.Authenticator.Client(context.Background(), token))
	lists, err := client.CurrentUsersPlaylists(context.Background())
	if err != nil {
		return nil, err
	}
	return lists, nil
}
func (c *ConcertifyCore) GetPlaylist(ctx context.Context, token *oauth2.Token, playlistId spotify.ID) (*spotify.FullPlaylist, error) {
	client := spotify.New(c.Authenticator.Client(context.Background(), token))
	playlist, err := client.GetPlaylist(context.Background(), playlistId)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}
