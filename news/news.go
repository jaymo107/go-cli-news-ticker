package news

var sources = []NewsGetter{
	NewHackerNews(),
	&DevToSource{},
}

type NewsGetter interface {
	GetNews() ([]Story, error)
}

func GetNews() ([]Story, error) {
	var stories []Story

	for _, source := range sources {
		if sourceStories, err := source.GetNews(); err == nil {
			stories = append(stories, sourceStories...)
		} else {
			return nil, err
		}
	}

	return stories, nil
}
