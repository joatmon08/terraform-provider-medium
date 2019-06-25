package medium

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStory(t *testing.T) {
	storyID := "c19b7de87dc4"
	authorID := "test"
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "{\"payload\":{\"value\": {\"id\": \"c19b7de87dc4\",\"versionId\": \"9f984568c5cf\",\"creatorId\": \"699911a5b2ab\",\"homeCollectionId\": \"\", \"title\": \"draft\",\"detectedLanguage\": \"\",\"latestVersion\": \"9f984568c5cf\",\"latestPublishedVersion\": \"\",\"hasUnpublishedEdits\": true}}}")
		}))
	defer ts.Close()

	endpoint := ReadEndpoint{
		Host: ts.URL,
	}
	story, err := endpoint.GetStory(authorID, storyID)
	if err != nil {
		t.Fatalf("unable to access story: %s", err)
	}
	if story.Payload.Value.ID != storyID {
		t.Fatalf("story is incorrect: %s", story.Payload.Value.ID)
	}
}
