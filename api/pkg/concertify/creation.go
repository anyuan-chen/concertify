package concertify

import (
	"context"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func (c *ConcertifyCore) CreateYoutubePlaylist(token *oauth2.Token) (string, error) {
	client, err := youtube.NewService(context.Background(), option.WithTokenSource(GoogleConfig.TokenSource(context.Background(), token)))
	if err != nil {
		return "", err
	}
	playlist := &youtube.Playlist{
		Snippet: &youtube.PlaylistSnippet{
			Title:       "Concertify",
			Description: "A playlist created by Concertify",
		},
		Status: &youtube.PlaylistStatus{
			PrivacyStatus: "private",
		},
	}
	createPlaylistCall := client.Playlists.Insert([]string{"snippet", "status"}, playlist)
	res, err := createPlaylistCall.Do()
	if err != nil {
		return "", err
	}
	return res.Id, nil
}

func (c *ConcertifyCore) InsertItemsIntoPlaylist(token *oauth2.Token, trackIds []string, id string) error {
	client, err := youtube.NewService(context.Background(), option.WithTokenSource(GoogleConfig.TokenSource(context.Background(), token)))
	if err != nil {
		return err
	}
	for _, trackId := range trackIds {
		item := &youtube.PlaylistItem{
			Snippet: &youtube.PlaylistItemSnippet{
				PlaylistId: id,
				ResourceId: &youtube.ResourceId{
					Kind:    "youtube#video",
					VideoId: trackId,
				},
			},
		}
		insertItemsCall := client.PlaylistItems.Insert([]string{"snippet"}, item)
		_, err = insertItemsCall.Do()
		if err != nil {
			return err
		}
	}
	return nil
}
