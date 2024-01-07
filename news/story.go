package news

import (
	"fmt"

	"github.com/jaymo107/go-cli-news-ticker/helpers"
)

type Story struct {
	Time  int
	Title string
	Url   string
}

type Stories []Story

func (s *Story) String() string {
	return fmt.Sprintf("(%s) - %s (%s)", s.getFormattedTime(), s.Title, s.Url)
}

func (s *Story) getFormattedTime() string {
	return helpers.GetTimeAgo(int64(s.Time))
}

func (s Stories) Len() int           { return len(s) }
func (s Stories) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Stories) Less(i, j int) bool { return s[i].Time > s[j].Time }
