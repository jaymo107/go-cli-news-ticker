package news

import (
	"encoding/json"

	"github.com/jaymo107/go-cli-news-ticker/helpers"
)

type DevToSource struct {
	topStoriesUrl string
	maxStories    int
}

func NewDevToSource() *DevToSource {
	return &DevToSource{
		topStoriesUrl: "https://dev.to/api/articles",
		maxStories:    10,
	}
}

func (s *DevToSource) GetNews() ([]Story, error) {
	if resp, err := helpers.GetJson(s.topStoriesUrl); err == nil {
		var stories []Story
		json.Unmarshal(resp, &stories)
		return stories, nil
	} else {
		return nil, err
	}
}
