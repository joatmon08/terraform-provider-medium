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
			"post_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"content_format": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"publish_status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePostCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	var publishStatus medium.PublishStatus

	switch status := d.Get("publish_status").(string); status {
	case medium.PublishStatusPublic:
		publishStatus = medium.PublishStatusPublic
	case medium.PublishStatusUnlisted:
		publishStatus = medium.PublishStatusUnlisted
	default:
		publishStatus = medium.PublishStatusDraft
	}

	options := medium.CreatePostOptions{
		UserID:        config.User.ID,
		Title:         d.Get("title").(string),
		Content:       d.Get("content").(string),
		ContentFormat: medium.ContentFormatMarkdown,
		PublishStatus: publishStatus,
	}

	post, err := config.Client.CreatePost(options)
	if err != nil {
		return err
	}
	d.SetId(post.ID)
	return resourcePostRead(d, m)
}

func resourcePostRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePostUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcePostRead(d, m)
}

func resourcePostDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
