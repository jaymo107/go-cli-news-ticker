package news

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"

	"github.com/jaymo107/go-cli-news-ticker/helpers"
)

type HackerNewsSource struct {
	topStoriesUrl string
	itemUrl       string
	maxStories    int
}

type HackerNewsSourceResponse struct {
	Time  int    `json:"time"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

func NewHackerNews() *HackerNewsSource {
	return &HackerNewsSource{
		topStoriesUrl: "https://hacker-news.firebaseio.com/v0/topstories.json",
		itemUrl:       "https://hacker-news.firebaseio.com/v0/item/%d.json",
		maxStories:    10,
	}
}

func (s *HackerNewsSource) GetNews() ([]Story, error) {
	storyIds, err := s.getTopStoryIds()

	if err != nil {
		return nil, err
	}

	var (
		stories []Story
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)

	for _, storyId := range storyIds {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			story, _ := s.getStory(id)
			mutex.Lock()
			stories = append(stories, *story)
			mutex.Unlock()
		}(storyId)
	}

	wg.Wait()

	return stories, nil
}

func (s *HackerNewsSource) getStory(storyId int) (*Story, error) {
	if resp, err := helpers.GetJson(fmt.Sprintf(s.itemUrl, storyId)); err == nil {
		var response HackerNewsSourceResponse
		json.Unmarshal(resp, &response)
		story := Story(response)
		return &story, nil
	} else {
		return nil, err
	}
}

func (s *HackerNewsSource) getTopStoryIds() ([]int, error) {
	if resp, err := helpers.GetJson(s.topStoriesUrl); err != nil {
		return nil, err
	} else {
		var topStoryIds []int
		json.Unmarshal(resp, &topStoryIds)
		storyIds := randomiseTopStories(topStoryIds, s.maxStories)

		return storyIds, nil
	}
}

func randomiseTopStories(ids []int, max int) []int {
	length := len(ids)
	randomIds := make([]int, max)

	for i := 0; i < len(randomIds); i++ {
		index := rand.Intn(length)
		randomIds[i] = ids[index]
	}

	return randomIds
}
