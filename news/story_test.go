package news

import (
	"fmt"
	"testing"
	"time"
)

func TestStoryString(t *testing.T) {
	story := Story{
		Time:  int(time.Now().Unix()),
		Title: "Test Story",
		Url:   "https://example.com/test",
	}

	expectedOutput := fmt.Sprintf("(%s) - %s (%s)", story.getFormattedTime(), story.Title, story.Url)
	result := story.String()

	if result != expectedOutput {
		t.Errorf("Expected %s, but got %s", expectedOutput, result)
	}
}
