package helpers

import (
	"testing"
	"time"
)

func TestGetTimeAgo(t *testing.T) {
	unixTime := time.Now().Add(-2*24*time.Hour - 3*time.Hour).Unix()
	expectedResult := "2 days ago"

	result := GetTimeAgo(unixTime)

	if result != expectedResult {
		t.Errorf("Expected '%s', but got '%s'", expectedResult, result)
	}

	unixTime = time.Now().Add(-1 * time.Hour).Unix()
	expectedResult = "1 hour ago"

	result = GetTimeAgo(unixTime)

	if result != expectedResult {
		t.Errorf("Expected '%s', but got '%s'", expectedResult, result)
	}

	unixTime = time.Now().Unix()
	expectedResult = "just now"

	result = GetTimeAgo(unixTime)

	if result != expectedResult {
		t.Errorf("Expected '%s', but got '%s'", expectedResult, result)
	}
}
