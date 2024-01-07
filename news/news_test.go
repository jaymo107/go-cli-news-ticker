package news

import (
	"sort"
	"testing"
)

type MockNewsGetter struct {
	Stories []Story
	Err     error
}

func (m MockNewsGetter) GetNews() ([]Story, error) {
	return m.Stories, m.Err
}

func TestGetNews(t *testing.T) {
	mockStories1 := []Story{{Time: 100, Title: "Mock Story 1", Url: "https://example.com/story1"}}
	mockStories2 := []Story{{Time: 200, Title: "Mock Story 2", Url: "https://example.com/story2"}}

	mockSource1 := MockNewsGetter{Stories: mockStories1, Err: nil}
	mockSource2 := MockNewsGetter{Stories: mockStories2, Err: nil}

	sources = []NewsGetter{mockSource1, mockSource2}

	result, err := GetNews()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedStories := append(mockStories1, mockStories2...)
	sort.Sort(Stories(expectedStories))

	if len(result) != 2 {
		t.Errorf("Expected %d stories, but got %d", len(expectedStories), len(result))
	}

	for i := range result {
		if result[i] != expectedStories[i] {
			t.Errorf("Mismatch in story at index %d. Expected %v, but got %v", i, expectedStories[i], result[i])
		}
	}
}
