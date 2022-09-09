package concertify

import (
	"context"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func (c *ConcertifyCore) CreateYoutubePlaylist() (string, error) {
	client, err := youtube.NewService(context.Background(), option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))
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

func (c *ConcertifyCore) InsertItemsIntoPlaylist(trackIds []string, id string) error {
	client, err := youtube.NewService(context.Background(), option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))
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
