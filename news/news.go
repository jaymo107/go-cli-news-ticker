package news

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jaymo107/go-cli-news-ticker/helpers"
)

const (
	topStoriesUrl = "https://hacker-news.firebaseio.com/v0/topstories.json"
	itemUrl       = "https://hacker-news.firebaseio.com/v0/item/%d.json"
)

func GetNews() ([]Story, error) {
	storyIds, err := getTopStoryIds()

	var (
		stories []Story
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)

	if err != nil {
		return nil, err
	}

	for _, storyId := range storyIds {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			story, _ := getStory(id)
			mutex.Lock()
			stories = append(stories, *story)
			mutex.Unlock()
		}(storyId)
	}

	wg.Wait()

	return stories, nil
}

func getStory(storyId int) (*Story, error) {
	resp, err := helpers.GetJson(fmt.Sprintf(itemUrl, storyId))

	if err != nil {
		return nil, err
	}

	var story Story
	json.Unmarshal(resp, &story)

	return &story, nil
}

func getTopStoryIds() ([]int, error) {
	resp, err := helpers.GetJson(topStoriesUrl)

	if err != nil {
		return nil, err
	}

	var topStoryIds []int
	json.Unmarshal(resp, &topStoryIds)
	storyIds := topStoryIds[:5]

	return storyIds, nil
}
