package medium

import (
	medium "github.com/medium/medium-sdk-go"
)

type PostBuilder struct {
	Options *medium.CreatePostOptions
}

func (b *PostBuilder) Build(userID string, title string, content string, publish_status string) {
	b.Options = &medium.CreatePostOptions{
		UserID:        userID,
		Title:         title,
		Content:       content,
		ContentFormat: medium.ContentFormatMarkdown,
		PublishStatus: retrievePublishStatus(publish_status),
	}
}

func (b *PostBuilder) Tags(tagsRaw []interface{}) {
	if len(tagsRaw) > 0 {
		tags := make([]string, len(tagsRaw))
		for i, tag := range tagsRaw {
			tags[i] = tag.(string)
		}
		b.Options.Tags = tags
	}
}

func retrievePublishStatus(status string) medium.PublishStatus {
	var publishStatus medium.PublishStatus
	switch s := status; s {
	case medium.PublishStatusPublic:
		publishStatus = medium.PublishStatusPublic
	case medium.PublishStatusUnlisted:
		publishStatus = medium.PublishStatusUnlisted
	case draftStatus:
		publishStatus = medium.PublishStatusDraft
	}
	return publishStatus
}
