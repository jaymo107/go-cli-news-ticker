package news

type DevToSource struct{}

var (
	topStoriesUrl = "https://hacker-news.firebaseio.com/v0/topstories.json"
	itemUrl       = "https://hacker-news.firebaseio.com/v0/item/%d.json"
	maxStories    = 10
)

func (s *DevToSource) GetNews() ([]Story, error) {
	return []Story{}, nil
}
