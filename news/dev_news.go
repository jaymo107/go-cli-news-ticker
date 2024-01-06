package news

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
	return []Story{}, nil
}
