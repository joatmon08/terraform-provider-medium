package medium

import (
	"github.com/hashicorp/terraform/helper/schema"
	medium "github.com/medium/medium-sdk-go"
)

func ResourcePost() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostCreate,
		Read:   resourcePostRead,
		Update: resourcePostUpdate,
		Delete: resourcePostDelete,

		Schema: map[string]*schema.Schema{
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"publish_status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"published_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"has_unpublished_edits": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"medium_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func retrievePublishStatus(status string) medium.PublishStatus {
	publishStatus := medium.PublishStatusDraft
	switch s := status; s {
	case medium.PublishStatusPublic:
		publishStatus = medium.PublishStatusPublic
	case medium.PublishStatusUnlisted:
		publishStatus = medium.PublishStatusUnlisted
	}
	return publishStatus
}

func resourcePostCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)

	options := medium.CreatePostOptions{
		UserID:        config.User.ID,
		Title:         d.Get("title").(string),
		Content:       d.Get("content").(string),
		ContentFormat: medium.ContentFormatMarkdown,
		PublishStatus: retrievePublishStatus(d.Get("publish_status").(string)),
	}

	post, err := config.Client.CreatePost(options)
	if err != nil {
		return err
	}
	d.SetId(post.ID)
	return resourcePostRead(d, m)
}

func resourcePostRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	story, err := config.ReadEndpoint.GetStory(config.User.ID, d.Id())
	if err != nil {
		return err
	}
	d.Set("user_id", config.User.ID)
	d.Set("title", story.Payload.Value.Title)
	d.Set("version", story.Payload.Value.LatestVersion)
	d.Set("created_at", story.Payload.Value.CreatedAt)
	d.Set("updated_at", story.Payload.Value.UpdatedAt)
	d.Set("published_at", story.Payload.Value.LatestPublishedAt)
	d.Set("revision", story.Payload.Value.LatestRev)
	d.Set("has_unpublished_edits", story.Payload.Value.HasUnpublishedEdits)
	d.Set("medium_url", story.Payload.Value.MediumURL)
	return nil
}

func resourcePostUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcePostRead(d, m)
}

func resourcePostDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
