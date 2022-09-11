package concertify

import (
	"context"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var GoogleConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_ID"),
	ClientSecret: os.Getenv("GOOGLE_SECRET"),
	Endpoint:     google.Endpoint,
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URI"),
	Scopes:       []string{"https://www.googleapis.com/auth/youtubepartner"},
}

func (c *ConcertifyCore) GetYoutubeVideoFromSpotify(query string, numberOfVideos int) (*youtube.SearchListResponse, error) {
	client, err := youtube.NewService(context.Background(), option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))
	if err != nil {
		return nil, err
	}
	searchListCall := client.Search.List([]string{"id", "snippet"}).Q(query + "live concert").MaxResults(int64(numberOfVideos))
	res, err := searchListCall.Do()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ConcertifyCore) GetYoutubeVideos(queries []string) ([]*youtube.SearchListResponse, error) {
	client, err := youtube.NewService(context.Background(), option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))
	if err != nil {
		return nil, err
	}
	var responses []*youtube.SearchListResponse
	for _, query := range queries {
		searchListCall := client.Search.List([]string{"id", "snippet"}).Q(query + "live concert").MaxResults(int64(1))
		res, err := searchListCall.Do()
		if err != nil {
			return nil, err
		}
		responses = append(responses, res)
	}
	return responses, nil
}

func (c *ConcertifyCore) GetVideoFromLink(ctx context.Context, token *oauth2.Token, videoUrl string) (*youtube.Video, error) {
	client, err := youtube.NewService(ctx, option.WithTokenSource(GoogleConfig.TokenSource(ctx, token)))
	if err != nil {
		return nil, err
	}
	videoId := videoUrl[len("https://www.youtube.com/watch?v="):]
	videoListCall := client.Videos.List([]string{"id", "snippet"}).Id(videoId)
	res, err := videoListCall.Do()
	if err != nil {
		return nil, err
	}
	return res.Items[0], nil
}
