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

func (s *Story) String() string {
	return fmt.Sprintf("(%s) - %s (%s)", s.getFormattedTime(), s.Title, s.Url)
}

func (s *Story) getFormattedTime() string {
	return helpers.GetTimeAgo(int64(s.Time))
}
