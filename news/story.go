package news

import (
	"fmt"

	"github.com/jaymo107/go-cli-news-ticker/helpers"
)

type Story struct {
	Time               int    `json:"time"`
	Title              string `json:"title"`
	Url                string `json:"url"`
	PublishedTimestamp string `json:"published_timestamp"`
}

func (s *Story) String() string {
	return fmt.Sprintf("(%s) - %s (%s)", s.getFormattedTime(), s.Title, s.Url)
}

func (s *Story) getFormattedTime() string {
	return helpers.GetTimeAgo(int64(s.Time))
}
