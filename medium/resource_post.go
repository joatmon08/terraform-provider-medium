package medium

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	medium "github.com/medium/medium-sdk-go"
)

const (
	draftStatus       = "draft"
	htmlContentFormat = "html"
)

var (
	allPublishStatuses = []string{
		draftStatus,
		medium.PublishStatusUnlisted,
		medium.PublishStatusPublic,
	}

	allContentFormats = []string{
		htmlContentFormat,
		medium.ContentFormatMarkdown,
	}
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
			"content_format": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(allContentFormats, false),
			},
			"publish_status": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(allPublishStatuses, false),
			},
			"tags": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
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

func resourcePostCreate(d *schema.ResourceData, m interface{}) error {
	var builder PostBuilder
	config := m.(*Config)
	userID := config.User.ID
	title := d.Get("title").(string)
	content := d.Get("content").(string)
	contentFormat := d.Get("content_format").(string)
	publishStatus := d.Get("publish_status").(string)
	tags := d.Get("tags").([]interface{})

	builder.BuildPost(userID, title, content, contentFormat, publishStatus)
	builder.Tags(tags)

	post, err := config.Client.CreatePost(*builder.PostOptions)
	if err != nil {
		return err
	}
	d.SetId(post.ID)
	return resourcePostRead(d, m)
}

func resourcePostRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	story, err := config.StoryEndpoint.GetStory(config.User.ID, d.Id())
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
	d.Set("tags", story.Payload.Value.Virtuals.Tags)
	return nil
}

func resourcePostUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcePostCreate(d, m)
}

func resourcePostDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
