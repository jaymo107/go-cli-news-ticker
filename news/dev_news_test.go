// devtosource_test.go
package news

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestDevToSource_GetNews(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockResponse := []DevToSourceResponse{
			{Title: "Test Story 1", Url: "https://example.com/story1", PublishedTimestamp: "2022-01-07T12:00:00Z"},
			{Title: "Test Story 2", Url: "https://example.com/story2", PublishedTimestamp: "2022-01-06T15:30:00Z"},
		}
		jsonResponse, _ := json.Marshal(mockResponse)
		w.Write(jsonResponse)
	}))

	defer server.Close()
	devToSource := &DevToSource{topStoriesUrl: server.URL, maxStories: 10}

	result, err := devToSource.GetNews()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedStories := []Story{
		{Time: int(time.Date(2022, 1, 7, 12, 0, 0, 0, time.UTC).Unix()), Title: "Test Story 1", Url: "https://example.com/story1"},
		{Time: int(time.Date(2022, 1, 6, 15, 30, 0, 0, time.UTC).Unix()), Title: "Test Story 2", Url: "https://example.com/story2"},
	}

	if !reflect.DeepEqual(result, expectedStories) {
		t.Errorf("Mismatch in the result. Expected %v, but got %v", expectedStories, result)
	}
}
