package helpers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJson(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"key": "value"}`))
	}))
	defer mockServer.Close()

	result, err := GetJson(mockServer.URL)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedResult := []byte(`{"key": "value"}`)

	if string(result) != string(expectedResult) {
		t.Errorf("Mismatch in the result. Expected %s, but got %s", expectedResult, result)
	}
}
