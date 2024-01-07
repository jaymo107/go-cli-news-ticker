package news

import "sort"

var sources = []NewsGetter{
	NewHackerNews(),
	NewDevToSource(),
}

type NewsGetter interface {
	GetNews() ([]Story, error)
}

func GetNews() ([]Story, error) {
	var stories Stories

	for _, source := range sources {
		if sourceStories, err := source.GetNews(); err == nil {
			stories = append(stories, sourceStories...)
		} else {
			return nil, err
		}
	}

	sort.Sort(stories)

	return stories, nil
}
