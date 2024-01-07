package news

import (
	"encoding/json"
	"time"

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
			parsedDate, err := response.getParsedDate()
			if err != nil {
				return nil, err
			}

			stories = append(stories, Story{
				Time:  parsedDate,
				Title: response.Title,
				Url:   response.Url,
			})
		}
		return stories, nil
	} else {
		return nil, err
	}
}

func (r *DevToSourceResponse) getParsedDate() (int, error) {
	format := "2006-01-02T15:04:05Z"
	if parsed, err := time.Parse(format, r.PublishedTimestamp); err == nil {
		return int(parsed.Unix()), nil
	} else {
		return 0, err
	}
}
