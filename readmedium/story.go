package readmedium

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	MediumURL = "https://medium.com"
	format    = "format=json"
)

type StoryEndpoint struct {
	Host string
}

type Story struct {
	Payload struct {
		Value struct {
			ID                     string `json:"id"`
			VersionID              string `json:"versionId"`
			CreatorID              string `json:"creatorId"`
			Title                  string `json:"title"`
			LatestVersion          string `json:"latestVersion"`
			LatestPublishedVersion string `json:"latestPublishedVersion"`
			HasUnpublishedEdits    bool   `json:"hasUnpublishedEdits"`
			LatestRev              int    `json:"latestRev"`
			CreatedAt              int64  `json:"createdAt"`
			UpdatedAt              int64  `json:"updatedAt"`
			AcceptedAt             int    `json:"acceptedAt"`
			FirstPublishedAt       int    `json:"firstPublishedAt"`
			LatestPublishedAt      int    `json:"latestPublishedAt"`
			Visibility             int    `json:"visibility"`
			License                int    `json:"license"`
			CanonicalURL           string `json:"canonicalUrl"`
			WebCanonicalURL        string `json:"webCanonicalUrl"`
			MediumURL              string `json:"mediumUrl"`
			Type                   string `json:"type"`
		} `json:"value"`
	} `json:"payload"`
}

func (r *StoryEndpoint) GetStory(author_id string, post_id string) (*Story, error) {
	var body []byte
	url := fmt.Sprintf("%s/%s/%s?%s", r.Host, author_id, post_id, format)
	req, err := http.NewRequest("GET", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "text/plain")
	req.Header.Add("Accept-Charset", "utf-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	storyBody := []byte(strings.ReplaceAll(string(response), "])}while(1);</x>", ""))
	if err != nil {
		return nil, err
	}

	var story Story
	if err := json.Unmarshal(storyBody, &story); err != nil {
		return nil, err
	}
	return &story, nil
}
