package news

import (
	"encoding/json"

	"github.com/jaymo107/go-cli-news-ticker/helpers"
)

type DevToSource struct {
	topStoriesUrl string
	maxStories    int
}

type DevToSourceResponse struct {
	Title              string `json:"title"`
	Url                string `json:"url"`
	PublishedTimestamp string `json:"published_timestamp"`
}

func NewDevToSource() *DevToSource {
	return &DevToSource{
		topStoriesUrl: "https://dev.to/api/articles",
		maxStories:    10,
	}
}

func (s *DevToSource) GetNews() ([]Story, error) {
	if resp, err := helpers.GetJson(s.topStoriesUrl); err == nil {
		var responses []DevToSourceResponse
		json.Unmarshal(resp, &responses)
		stories := []Story{}
		for _, response := range responses {
			stories = append(stories, Story{
				Time:  0, // TODO: Parse this properly
				Title: response.Title,
				Url:   response.Url,
			})
		}
		return stories, nil
	} else {
		return nil, err
	}
}
